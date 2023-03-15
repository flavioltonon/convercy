package dto

import (
	"net/http"
	"strconv"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gorilla/mux"
)

type ConvertCurrencyRequest struct {
	Amount float64
	Code   string
}

func (r ConvertCurrencyRequest) Validate() error {
	return ozzo.ValidateStruct(&r,
		ozzo.Field(&r.Amount, ozzo.Required),
		ozzo.Field(&r.Code, ozzo.Required),
	)
}

func ParseConvertCurrencyRequest(r *http.Request) (ConvertCurrencyRequest, error) {
	pathParams := mux.Vars(r)

	currencyAmount, err := strconv.ParseFloat(pathParams["currency_amount"], 64)
	if err != nil {
		return ConvertCurrencyRequest{}, err
	}

	request := ConvertCurrencyRequest{
		Amount: currencyAmount,
		Code:   pathParams["currency_code"],
	}

	if err := request.Validate(); err != nil {
		return ConvertCurrencyRequest{}, err
	}

	return request, nil
}
