package models

type Account struct {
	Id             int64
	DocumentNumber string
	AvailableLimit float64
	CurrentLimit   float64

	AccountLimit
}

type AccountLimit struct {
	AvailableLimit float64
	CurrentLimit   float64
}
