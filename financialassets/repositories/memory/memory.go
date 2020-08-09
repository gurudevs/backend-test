package memory

import (
	"sync"

	"github.com/ferkze/backend-test/financialassets/model"
	"github.com/ferkze/backend-test/financialassets/repositories"
)

// FinancialAssetRepository repositório de ativos financeiros na memória
type FinancialAssetRepository struct {
	mu    *sync.Mutex
	assets map[string]*FinancialAsset
}

// FinancialAsset estrutura do dado de ativo financeiro na memória
type FinancialAsset struct {

	Ticker string
	Company string
	Close float32
	Open float32
	Price float32
	PctVariation float32
	PriceVariation float32

}

// NewFinancialAssetRepository instancia do financialAssetRepository
func NewFinancialAssetRepository() repositories.FinancialAssetRepository {
	return &FinancialAssetRepository{
			mu:    &sync.Mutex{},
			assets: map[string]*FinancialAsset{},
	}
}

// FindAll busca todos os ativos financeiros no repositório
func (r *FinancialAssetRepository) FindAll() ([]*model.FinancialAsset, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	assets := make([]*model.FinancialAsset, len(r.assets))
	i := 0
	for _, asset := range r.assets {
		assets[i] = model.NewFinancialAsset(asset.Ticker)
		i++
	}
	return assets, nil
}

// Add adiciona um ativo financeiro à memória
func (r *FinancialAssetRepository) Add(asset *model.FinancialAsset) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.assets[asset.GetTicker()] = &FinancialAsset{
		Ticker: asset.Ticker,
		Company: asset.Company,
		Close: asset.Close,
		Price: asset.Price,
		Open: asset.Open,
		PctVariation: asset.PctVariation,
		PriceVariation: asset.PriceVariation,
	}
	return nil
}