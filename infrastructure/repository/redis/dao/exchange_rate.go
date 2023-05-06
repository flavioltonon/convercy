package dao

import "encoding/json"

type CurrencyExchangeRates struct {
	CurrencyCode  string
	ExchangeRates []ExchangeRate
}

func (dao CurrencyExchangeRates) MarshalBinary() ([]byte, error) {
	return json.Marshal(dao)
}

func (dao *CurrencyExchangeRates) UnmarshalBinary(b []byte) error {
	return json.Unmarshal(b, dao)
}

type ExchangeRate struct {
	Rate float64
	Unit ExchangeRateUnit
}

func (dao ExchangeRate) MarshalBinary() ([]byte, error) {
	return json.Marshal(dao)
}

func (dao *ExchangeRate) UnmarshalBinary(b []byte) error {
	return json.Unmarshal(b, dao)
}

type ExchangeRateUnit struct {
	Base   string
	Target string
}

func (dao ExchangeRateUnit) MarshalBinary() ([]byte, error) {
	return json.Marshal(dao)
}

func (dao *ExchangeRateUnit) UnmarshalBinary(b []byte) error {
	return json.Unmarshal(b, dao)
}
