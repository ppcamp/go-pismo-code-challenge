package dtos

type CreateAccount struct {
	DocumentNumber string `json:"document_number"`
}

type Account struct {
	Id             int64  `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}
