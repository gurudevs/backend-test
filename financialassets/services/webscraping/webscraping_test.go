package webscraping

import (
	"os"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestGetIbovespaAssetTickers(t *testing.T) {
	t.Run("Deve fazer a busca de tickers da composição do ibovespa", func (tt *testing.T) {
		srv := NewFinancialAssetScraperService()
	
		tickers, err := srv.GetIbovespaAssetTickers()
		if err != nil {
			tt.Fatalf("Erro não esperado ao fazer scraping de tickers da composição do ibovespa: %s\n", err.Error())
		}
		if len(tickers) == 0 {
			tt.Fatalf("Esperado que o scraping de tickers da composição do ibovespa retornasse resultados\n")
		}
		tt.Logf("Recebido %d tickers com sucesso\n", len(tickers))
	})
}

func TestGetAssetData(t *testing.T) {
	t.Run("Deve fazer a busca de dados do ativo", func (tt *testing.T) {
		file, err := os.Open("docs/ITUB3_page.html")
		if err != nil {
			tt.Fatalf("Falha na leitura de arquivo: %s\n", err.Error())
		}
		doc, err := goquery.NewDocumentFromReader(file)
		if err != nil {
			tt.Fatalf("Falha no parse do documento HTML: %s\n", err.Error())
		}

		asset, err := scrapeAsset(doc)
		if err != nil {
			tt.Fatalf("Falha no scrape dos dados no documento: %s\n", err.Error())
		}

		if asset.Ticker != "ITUB3" {
			tt.Errorf("Erro no ticker do ativo, esperado: %s, obtido: %s\n", "ITUB3", asset.Ticker)
		}

		tt.Logf("Encontrado dados do ativo da Companhia %s\n", asset.Company)

	})
}