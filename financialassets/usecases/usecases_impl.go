package usecases

import (
	"log"
	"math"
	"sort"

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
	sort.SliceStable(assets, func(i, j int) bool {
    return math.Abs(float64(assets[i].PctVariation)) > math.Abs(float64(assets[j].PctVariation))
	})
	return assets, nil
}

func (uc *financialAssetsUsecasesImpl) PopulateAssets() (error) {
	tickers, err := uc.srv.GetIbovespaAssetTickers()
	if err != nil {
		return err
	}
	tickers = tickers[0:10]
	assetCh := make(chan *model.FinancialAsset, len(tickers))
	errCh := make(chan error)
	for _, ticker := range tickers {
		go uc.srv.GetAssetDataCh(ticker, assetCh, errCh)
	}
	assets := make([]*model.FinancialAsset, len(tickers))
	i := 0
	for asset := range assetCh {
		select {
		case err := <- errCh:
			log.Println("Erro recebido no canal, abortando busca de ativos...")
			close(assetCh)
			return err
		default:
			log.Printf("O ativo %d de %d do Ibovespa está sendo guardado\n", i, len(tickers))
			assets[i] = asset
			if i == len(tickers)-1 {
				close(assetCh)
				close(errCh)
				break
			}
			i++
		}
	}
	return uc.repo.Setup(assets)
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