package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"convercy/application"
	"convercy/application/config"
	"convercy/application/http/controllers/backoffice"
	"convercy/application/http/controllers/user"
	"convercy/application/http/middleware"
	
	domainServices "convercy/domain/services"
	"convercy/domain/valueobject"
	"convercy/infrastructure/logging/zap"
	"convercy/infrastructure/repository/mongodb"
	mongodbMappers "convercy/infrastructure/repository/mongodb/mappers"
	"convercy/infrastructure/repository/openexchangerates"
	"convercy/infrastructure/repository/redis"
	redisMappers "convercy/infrastructure/repository/redis/mappers"
	"convercy/shared/logging"

	"github.com/gorilla/mux"
)

func main() {
	config, err := config.New()
	if err != nil {
		panic(err)
	}

	logger, err := zap.NewLogger()
	if err != nil {
		panic(err)
	}

	repository, err := mongodb.NewRepository(&mongodb.Options{
		Database: mongodb.DatabaseOptions{
			Name: "convercy",
			URI:  config.Database.URI,
		},
	})
	if err != nil {
		logger.Fatal("failed to create repository", logging.Error(err))
	}

	if err := repository.Connect(); err != nil {
		logger.Fatal("failed to connect to repository", logging.Error(err))
	}

	defer repository.Disconnect()

	cache, err := redis.NewCache(&redis.Options{
		Address:           config.Cache.Address,
		ConnectionTimeout: config.Cache.ConnectionTimeout,
	})
	if err != nil {
		logger.Fatal("failed to create cache", logging.Error(err))
	}

	var (
		baselineCurrencyCode, _              = valueobject.NewCurrencyCode("USD")
		openExchangeRatesClient              = openexchangerates.NewClient(http.DefaultClient, config.OpenExchangeRates.AppID, config.OpenExchangeRates.BaseURL)
		currenciesRepository                 = openexchangerates.NewCurrenciesRepository(openExchangeRatesClient)
		exchangeRateUnitsMapper              = redisMappers.NewExchangeRateUnitsMapper()
		exchangeRatesMapper                  = redisMappers.NewExchangeRatesMapper(exchangeRateUnitsMapper)
		currencyExchangeRatesMapper          = redisMappers.NewCurrencyExchangeRatesMapper(exchangeRatesMapper)
		currencyExchangeRatesCache           = redis.NewCurrencyExchangeRatesCache(currencyExchangeRatesMapper, cache, config.Cache.CurrencyExchangeRates.TTL)
		currencyExchangeRatesRepository      = openexchangerates.NewCurrencyExchangeRatesRepository(openExchangeRatesClient)
		currencyConversionDomainService      = domainServices.NewCurrencyConversionService()
		currencyMapper                       = mongodbMappers.NewCurrencyMapper()
		registeredCurrenciesMapper           = mongodbMappers.NewRegisteredCurrenciesMapper(currencyMapper)
		registeredCurrenciesRepository       = mongodb.NewRegisteredCurrenciesRepository(registeredCurrenciesMapper, repository)
		currencyConversionApplicationService = application.NewCurrencyConversionService(
			baselineCurrencyCode,
			currenciesRepository,
			currencyConversionDomainService,
			currencyExchangeRatesCache,
			currencyExchangeRatesRepository,
			registeredCurrenciesRepository,
		)
		currencyRegistrationApplicationService = application.NewCurrencyRegistrationService(currenciesRepository, registeredCurrenciesRepository)
		backofficeCurrencyController           = backoffice.NewCurrencyController(currencyRegistrationApplicationService, logger)
		userCurrencyController                 = user.NewCurrencyController(currencyConversionApplicationService, logger)
	)

	// Setup router
	router := mux.NewRouter()
	router.HandleFunc("/api/convert/{currency_code}/{currency_amount}", userCurrencyController.ConvertCurrency).Methods(http.MethodGet)
	router.HandleFunc("/api/backoffice/currencies", backofficeCurrencyController.ListRegisteredCurrencies).Methods(http.MethodGet)
	router.HandleFunc("/api/backoffice/currencies", backofficeCurrencyController.RegisterCurrency).Methods(http.MethodPost)
	router.HandleFunc("/api/backoffice/currencies/{currency_id}", backofficeCurrencyController.UnregisterCurrency).Methods(http.MethodDelete)
	router.Use(middleware.Log(logger))

	server := &http.Server{
		Addr:    config.Server.Address,
		Handler: router,
	}

	go func() {
		logger.Info("listening and serving", logging.String("server_address", server.Addr))
		server.ListenAndServe()
	}()

	interrupt := make(chan os.Signal, 1)

	signal.Notify(
		interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
	)

	<-interrupt

	log.Println("gracefully shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
