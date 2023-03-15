package dao

type Currency struct {
	ID   string `bson:"_id"`
	Code string `bson:"code"`
}
