package repository

import "github.com/ferkze/backend-test/financialassets/domain/model"

// FinancialAssetRepository interface/contrato de repositorio de ativos financeiros
type FinancialAssetRepository interface {
	
	FindAll() ([]*model.FinancialAsset, error)

	FindByTicker(ticker string) (*model.FinancialAsset, error)
	
}