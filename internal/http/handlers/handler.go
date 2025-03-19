package handlers

import (
	"github.com/ppcamp/go-pismo-code-challenge/internal/services"
)

// Handler is just a shortcut struct, working as an internal dto, to avoid
// passing a huge amount of args to HTTP server startup function.
type Handler struct {
	Account services.Account

	Transaction services.Transaction
}
