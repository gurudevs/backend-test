package model

// FinancialAsset Ativo financeiro
type FinancialAsset struct {
	Ticker         string  `json:"ticker"`
	Company        string  `json:"company"`
	Close          float32 `json:"close"`
	Open           float32 `json:"open"`
	Price          float32 `json:"price"`
	PctVariation   float32 `json:"pctVariation"`
	PriceVariation float32 `json:"priceVariation"`
}

// NewFinancialAsset instanciar novo ativo financeiro
func NewFinancialAsset(ticker, company string, close, open, price, pctVariation, priceVariation float32) *FinancialAsset {
	return &FinancialAsset{
		Ticker: ticker,
		Company: company,
		Open: open,
		Close: close,
		Price: price,
		PctVariation: pctVariation,
		PriceVariation: priceVariation,
	}
}

// GetTicker retorna o ticker do ativo financeiro
func (fa *FinancialAsset) GetTicker() string {
	return fa.Ticker
}
