package controllers

import (
	"errors"
	"net/http"

	"convercy/application/dto"
	"convercy/application/services"
	"convercy/domain"
	"convercy/infrastructure/response"
)

type CurrencyController struct {
	currencyConversionService   *services.CurrencyConversionService
	currencyRegistrationService *services.CurrencyRegistrationService
}

func NewCurrencyController(
	currencyConversionService *services.CurrencyConversionService,
	currencyRegistrationService *services.CurrencyRegistrationService) *CurrencyController {
	return &CurrencyController{
		currencyConversionService:   currencyConversionService,
		currencyRegistrationService: currencyRegistrationService,
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
			response.JSON(w, response.InternalServerError(err))
		}

		return
	}

	response.JSON(w, response.NoContent())
}

func (c *CurrencyController) ListRegisteredCurrencies(w http.ResponseWriter, r *http.Request) {
	result, err := c.currencyRegistrationService.ListRegisteredCurrencies()
	if err != nil {
		response.JSON(w, response.InternalServerError(err))
		return
	}

	response.JSON(w, response.OK(result))
}

func (c *CurrencyController) ConvertCurrency(w http.ResponseWriter, r *http.Request) {
	request, err := dto.ParseConvertCurrencyRequest(r)
	if err != nil {
		response.JSON(w, response.BadRequest(err))
		return
	}

	result, err := c.currencyConversionService.ConvertCurrency(request)
	if err != nil {
		switch {
		case errors.As(err, new(domain.ErrValidationFailure)):
			response.JSON(w, response.BadRequest(err))
		case errors.As(err, new(domain.ErrNotFound)):
			response.JSON(w, response.NotFound(err))
		default:
			response.JSON(w, response.InternalServerError(err))
		}

		return
	}

	response.JSON(w, response.OK(result))
}
