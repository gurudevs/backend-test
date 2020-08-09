package memory

import (
	"sync"
	"time"

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
	UpdatedAt time.Time

}

// NewFinancialAssetRepository instancia do financialAssetRepository
func NewFinancialAssetRepository() repositories.FinancialAssetRepository {
	return &FinancialAssetRepository{
			mu:    &sync.Mutex{},
			assets: map[string]*FinancialAsset{},
	}
}

// Setup inicializa as chaves de tickers dos ativos financeiros
func (r *FinancialAssetRepository) Setup(assets []*model.FinancialAsset) (error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, asset := range assets {
		r.assets[asset.GetTicker()] = &FinancialAsset{
			Company: asset.Company,
			Close: asset.Close,
			Open: asset.Open,
			PctVariation: asset.PctVariation,
			PriceVariation: asset.PriceVariation,
			Price: asset.Price,
			Ticker: asset.Ticker,
			UpdatedAt: time.Now(),
		}
	}

	return nil
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

// Set adiciona um ativo financeiro à memória
func (r *FinancialAssetRepository) Set(asset *model.FinancialAsset) error {
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
		UpdatedAt: asset.GetUpdatedAt(),
	}
	return nil
}