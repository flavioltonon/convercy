package backoffice

import (
	"errors"
	"net/http"

	"convercy/application/dto"
	"convercy/application/services"
	"convercy/domain"
	"convercy/infrastructure/response"
	"convercy/shared/logging"
)

type CurrencyController struct {
	currencyRegistrationService *services.CurrencyRegistrationService
	logger                      logging.Logger
}

func NewCurrencyController(currencyRegistrationService *services.CurrencyRegistrationService, logger logging.Logger) *CurrencyController {
	return &CurrencyController{
		currencyRegistrationService: currencyRegistrationService,
		logger:                      logger,
	}
}

func (c *CurrencyController) RegisterCurrency(w http.ResponseWriter, r *http.Request) {
	request, err := dto.ParseRegisterCurrencyRequest(r)
	if err != nil {
		response.JSON(w, response.BadRequest(err))
		return
	}

	result, err := c.currencyRegistrationService.RegisterCurrency(request)
	if err != nil {
		switch {
		case errors.As(err, new(domain.ErrValidationFailure)), errors.As(err, new(domain.ErrAlreadyExists)):
			response.JSON(w, response.BadRequest(err))
		default:
			c.logger.Error("failed to register currency", logging.Error(err))
			response.JSON(w, response.InternalServerError(err))
		}

		return
	}

	response.JSON(w, response.Created(result))
}

func (c *CurrencyController) UnregisterCurrency(w http.ResponseWriter, r *http.Request) {
	request, err := dto.ParseUnregisterCurrencyRequest(r)
	if err != nil {
		response.JSON(w, response.BadRequest(err))
		return
	}

	if err := c.currencyRegistrationService.UnregisterCurrency(request); err != nil {
		switch {
		case errors.As(err, new(domain.ErrValidationFailure)), errors.As(err, new(domain.ErrAlreadyExists)):
			response.JSON(w, response.BadRequest(err))
		case errors.As(err, new(domain.ErrNotFound)):
			response.JSON(w, response.NotFound(err))
		default:
			c.logger.Error("failed to unregister currency", logging.Error(err))
			response.JSON(w, response.InternalServerError(err))
		}

		return
	}

	response.JSON(w, response.NoContent())
}

func (c *CurrencyController) ListRegisteredCurrencies(w http.ResponseWriter, r *http.Request) {
	result, err := c.currencyRegistrationService.ListRegisteredCurrencies()
	if err != nil {
		c.logger.Error("failed to list registered currencies", logging.Error(err))
		response.JSON(w, response.InternalServerError(err))
		return
	}

	response.JSON(w, response.OK(result))
}
