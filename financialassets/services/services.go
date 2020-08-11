package services

import (
	"github.com/ferkze/backend-test/financialassets/model"
)

// FinancialAssetsServices serviços de ativos financeiros
type FinancialAssetsServices interface {

	GetIbovespaAssetTickers() ([]string, error)

	GetAssetData(string) (*model.FinancialAsset, error)

	GetAssetDataCh(string, chan *model.FinancialAsset, chan error)

}