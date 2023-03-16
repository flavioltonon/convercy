package user

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
	currencyConversionService *services.CurrencyConversionService
	logger                    logging.Logger
}

func NewCurrencyController(currencyConversionService *services.CurrencyConversionService, logger logging.Logger) *CurrencyController {
	return &CurrencyController{
		currencyConversionService: currencyConversionService,
		logger:                    logger,
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
			c.logger.Error("failed to convert currency", logging.Error(err))
			response.JSON(w, response.InternalServerError(err))
		}

		return
	}

	response.JSON(w, response.OK(result))
}
