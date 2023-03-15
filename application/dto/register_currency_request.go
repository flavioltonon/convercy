package dto

import (
	"encoding/json"
	"net/http"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
)

type RegisterCurrencyRequest struct {
	Code string `json:"code"`
}

func (r RegisterCurrencyRequest) Validate() error {
	return ozzo.ValidateStruct(&r,
		ozzo.Field(&r.Code, ozzo.Required),
	)
}

func ParseRegisterCurrencyRequest(r *http.Request) (RegisterCurrencyRequest, error) {
	var request RegisterCurrencyRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return RegisterCurrencyRequest{}, err
	}

	if err := request.Validate(); err != nil {
		return RegisterCurrencyRequest{}, err
	}

	return request, nil
}
