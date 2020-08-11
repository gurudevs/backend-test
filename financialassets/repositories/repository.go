package repositories

import "github.com/ferkze/backend-test/financialassets/model"

// FinancialAssetRepository interface/contrato de repositorio de ativos financeiros
type FinancialAssetRepository interface {
	
	Setup([]*model.FinancialAsset) (error)
	FindAll() ([]*model.FinancialAsset, error)
	Set(*model.FinancialAsset) error

	// FindByTicker(ticker string) (*model.FinancialAsset, error)
	
}