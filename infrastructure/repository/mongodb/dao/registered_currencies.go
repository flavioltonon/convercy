package dao

type RegisteredCurrencies struct {
	ClientID   string     `bson:"client_id"`
	Currencies []Currency `bson:"currencies"`
}
