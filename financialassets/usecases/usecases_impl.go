package usecases

import (
	"github.com/ferkze/backend-test/financialassets/model"
	"github.com/ferkze/backend-test/financialassets/repositories"
	"github.com/ferkze/backend-test/financialassets/services"
)

// financialAssetsUsecaseImpl estrutura do contrado de implementação do caso de uso
type financialAssetsUsecasesImpl struct {
	srv services.FinancialAssetsServices
	repo repositories.FinancialAssetRepository
}

// NewFinancialAssetsUsecases retorna instancia da implementação do caso de uso
func NewFinancialAssetsUsecases(srv services.FinancialAssetsServices, repo repositories.FinancialAssetRepository) FinancialAssetsUsecases {
	return &financialAssetsUsecasesImpl{
		srv: srv,
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

func (uc *financialAssetsUsecasesImpl) InitializeAssetQuotations() (error) {
	tickers, err := uc.srv.GetIbovespaAssetTickers()
	if err != nil {
		return err
	}
	for _, ticker := range tickers {
		asset, err := uc.srv.GetAssetData(ticker)
		if err != nil {
			return err
		}
		uc.repo.Set(asset)
	}
	return nil
}

// GetAssetsOrderedByVariation Seleciona todos os ativos ordenados por variação descendente
func (uc *financialAssetsUsecasesImpl) RefreshAssetQuotations() (error) {
	assets, err := uc.repo.FindAll()
	if err != nil {
		return err
	}
	for _, asset := range assets {
		asset, err := uc.srv.GetAssetData(asset.GetTicker())
		if err != nil {
			return err
		}
		uc.repo.Set(asset)
	}
	return nil
}