package usecases

import (
	"github.com/ferkze/backend-test/financialassets/model"
	"github.com/ferkze/backend-test/financialassets/repositories"
)

// financialAssetsUsecaseImpl estrutura do contrado de implementação do caso de uso
type financialAssetsUsecasesImpl struct {
	repo repositories.FinancialAssetRepository
}

// NewFinancialAssetsUsecases retorna instancia da implementação do caso de uso
func NewFinancialAssetsUsecases(repo repositories.FinancialAssetRepository) FinancialAssetsUsecases {
	return &financialAssetsUsecasesImpl{
		repo: repo,
	}
}

// GetAssetsOrderedByVariation Seleciona todos os ativos ordenados por variação descendente
func (uc *financialAssetsUsecasesImpl) GetAssetsOrderedByVariation() ([]*model.FinancialAsset, error) {
	assets, err := uc.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return assets, nil
}