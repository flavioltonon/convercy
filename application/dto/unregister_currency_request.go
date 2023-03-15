package dto

import (
	"net/http"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gorilla/mux"
)

type UnregisterCurrencyRequest struct {
	CurrencyID string `json:"currency_id"`
}

func (r UnregisterCurrencyRequest) Validate() error {
	return ozzo.ValidateStruct(&r,
		ozzo.Field(&r.CurrencyID, ozzo.Required),
	)
}

func ParseUnregisterCurrencyRequest(r *http.Request) (UnregisterCurrencyRequest, error) {
	vars := mux.Vars(r)

	request := UnregisterCurrencyRequest{
		CurrencyID: vars["currency_id"],
	}

	if err := request.Validate(); err != nil {
		return UnregisterCurrencyRequest{}, err
	}

	return request, nil
}
