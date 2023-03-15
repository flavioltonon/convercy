package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"convercy/application/config"
	"convercy/application/http/controllers"
	"convercy/application/http/middleware"
	applicationServices "convercy/application/services"
	domainServices "convercy/domain/services"
	"convercy/infrastructure/logging/zap"
	"convercy/infrastructure/openexchangerates"
	"convercy/infrastructure/repository/mongodb"
	"convercy/infrastructure/repository/mongodb/mappers"
	"convercy/shared/logging"

	"github.com/gorilla/mux"
)

func main() {
	config, err := config.FromFile("config.yaml")
	if err != nil {
		panic(err)
	}

	logger, err := zap.NewLogger()
	if err != nil {
		panic(err)
	}

	repository, err := mongodb.NewRepository(&mongodb.Options{
		Database: mongodb.DatabaseOptions{
			Name: "gruposbf",
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

	var (
		openExchangeRatesClient = openexchangerates.NewClient(
			http.DefaultClient,
			config.OpenExchangeRates.AppID,
			config.OpenExchangeRates.BaseURL,
		)
		currenciesService               = openexchangerates.NewCurrenciesService(openExchangeRatesClient)
		exchangeRatesService            = openexchangerates.NewExchangeRatesService(openExchangeRatesClient)
		currencyCodeValidationService   = domainServices.NewCurrencyCodeValidationService(currenciesService)
		currencyExchangeRatesService    = domainServices.NewCurrencyExchangeRatesService(exchangeRatesService)
		currencyConversionDomainService = domainServices.NewCurrencyConversionService()
		currencyMapper                  = mappers.NewCurrencyMapper()
		registeredCurrenciesMapper      = mappers.NewRegisteredCurrenciesMapper(currencyMapper)
		currenciesRepository            = mongodb.NewCurrenciesRepository(
			registeredCurrenciesMapper,
			repository,
		)
		currencyConversionApplicationService = applicationServices.NewCurrencyConversionService(
			currencyCodeValidationService,
			currencyConversionDomainService,
			currenciesRepository,
			currencyExchangeRatesService,
		)
		currencyRegistrationApplicationService = applicationServices.NewCurrencyRegistrationService(currencyCodeValidationService, currenciesRepository)
		currencyController                     = controllers.NewCurrencyController(
			currencyConversionApplicationService,
			currencyRegistrationApplicationService,
		)
	)

	// Setup router
	router := mux.NewRouter()
	router.HandleFunc("/api/convert/{currency_code}/{currency_amount}", currencyController.ConvertCurrency).Methods(http.MethodGet)
	router.HandleFunc("/api/backoffice/currencies", currencyController.ListRegisteredCurrencies).Methods(http.MethodGet)
	router.HandleFunc("/api/backoffice/currencies", currencyController.RegisterCurrency).Methods(http.MethodPost)
	router.HandleFunc("/api/backoffice/currencies/{currency_id}", currencyController.UnregisterCurrency).Methods(http.MethodDelete)
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
