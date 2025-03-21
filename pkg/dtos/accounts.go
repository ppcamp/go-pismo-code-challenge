package dtos

type ChangeAccountLimit struct {
	NewLimit float64 `json:"new_limit"`
}

type CreateAccount struct {
	DocumentNumber string `json:"document_number"`
}

type Account struct {
	Id             int64  `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}

type AccountLimits struct {
	CurrentLimit   float64 `json:"current_limit"`
	AvailableLimit float64 `json:"available_limit"`
}
