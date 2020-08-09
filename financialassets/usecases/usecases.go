package usecases

import (
	"github.com/ferkze/backend-test/financialassets/model"
)

// FinancialAssetsUsecases casos de uso do serviço
type FinancialAssetsUsecases interface {
	GetAssetsOrderedByVariation() ([]*model.FinancialAsset, error)

	InitializeAssetQuotations() (error)
	RefreshAssetQuotations() (error)
}

// FinancialAsset Estrutura DAO do ativo financeiro
type FinancialAsset struct {

	Ticker string
	Company string
	Close float32
	Open float32
	Price float32
	PctVariation float32
	PriceVariation float32

}

func toFinancialAsset(assets []*model.FinancialAsset) ([]*FinancialAsset) {
	res := make([]*FinancialAsset, len(assets))
	for i, asset := range assets {
		res[i] = &FinancialAsset{
			Ticker: asset.GetTicker(),
		}
	}
	return res
}