package model

// FinancialAsset Ativo financeiro
type FinancialAsset struct {
	Ticker string
	Company string
	Close float32
	Open float32
	Price float32
	PctVariation float32
	PriceVariation float32

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