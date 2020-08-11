package controllers

import "net/http"

// FinancialAssetsController interface de controller de ativos financeiros
type FinancialAssetsController interface {
	GetAssetsOrderedByVariation(w http.ResponseWriter, r *http.Request)
}