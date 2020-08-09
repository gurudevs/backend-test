package webscraping

import (
	"github.com/ferkze/backend-test/financialassets/model"
	"github.com/ferkze/backend-test/financialassets/services"
)

type WebScrapingService struct {

}

// NewFinancialAssetScraperService instancia do financialAssetsServices
func NewFinancialAssetScraperService() services.FinancialAssetsServices {
	return &WebScrapingService{}
}

// GetIbovespaAssetTickers retorna os tickers dos ativos que compõem o ibovespa
func (s *WebScrapingService) GetIbovespaAssetTickers() ([]string, error) {
	return []string{}, nil
}
// GetAssetData retorna o resultado do scraping dado o ticker do ativo
func (s *WebScrapingService) GetAssetData(ticker string) (*model.FinancialAsset, error) {
	return nil, nil
}