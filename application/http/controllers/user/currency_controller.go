package user

import (
	"errors"
	"net/http"

	"convercy/application/dto"
	"convercy/application/services"
	"convercy/domain"
	"convercy/infrastructure/response"
)

type CurrencyController struct {
	currencyConversionService *services.CurrencyConversionService
}

func NewCurrencyController(currencyConversionService *services.CurrencyConversionService) *CurrencyController {
	return &CurrencyController{
		currencyConversionService: currencyConversionService,
	}
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
