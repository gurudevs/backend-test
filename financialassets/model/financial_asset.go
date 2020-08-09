package model

import "time"

// FinancialAsset Ativo financeiro
type FinancialAsset struct {
	Ticker         string  `json:"ticker"`
	Company        string  `json:"company"`
	Close          float32 `json:"close"`
	Open           float32 `json:"open"`
	Price          float32 `json:"price"`
	PctVariation   float32 `json:"pctVariation"`
	PriceVariation float32 `json:"priceVariation"`
	updatedAt      time.Time
}

// NewFinancialAsset instanciar novo ativo financeiro
func NewFinancialAsset(ticker string) *FinancialAsset {
	return &FinancialAsset{
		Ticker: ticker,
	}
}

// GetTicker retorna o ticker do ativo financeiro
func (fa *FinancialAsset) GetTicker() string {
	return fa.Ticker
}

// GetUpdatedAt retorna o objeto de tempo de sua ultima atualização
func (fa *FinancialAsset) GetUpdatedAt() time.Time {
	return fa.updatedAt
}
