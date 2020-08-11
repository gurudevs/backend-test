package webscraping

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/ferkze/backend-test/financialassets/model"
	"github.com/ferkze/backend-test/financialassets/services"
)

const (
	tickerSelector = "#SymbolTitle > div.inline.heading > div > div:nth-child(2) > h1.symbol-h1 > strong"
	currentQuotationSelector = "#quoteElementPiece1"
	companyNameSelector = "#quote_top > div:nth-child(6) > table > tbody > tr.odd > td:nth-child(1)"
	priceVariationSelector = "#quoteElementPiece8"
	pctVariationSelector = "#quoteElementPiece9"
	openQuotationSelector = "#quoteElementPiece13"
	closeQuotationSelector = "#quoteElementPiece14"
)

type WebScrapingService struct {

}

// NewFinancialAssetScraperService instancia do financialAssetsServices
func NewFinancialAssetScraperService() services.FinancialAssetsServices {
	return &WebScrapingService{}
}

// GetIbovespaAssetTickers retorna os tickers dos ativos que compõem o ibovespa
func (s *WebScrapingService) GetIbovespaAssetTickers() ([]string, error) {
	tickers := []string{}
	
	res, err := http.Get("http://bvmf.bmfbovespa.com.br/indices/ResumoCarteiraTeorica.aspx?Indice=IBOV&idioma=pt-br")
  if err != nil {
    log.Println(err)
		return tickers, err
  }
  defer res.Body.Close()
  if res.StatusCode != 200 {
    log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		return tickers, fmt.Errorf("Unexpected page response status %d", res.StatusCode)
  }

  // Load the HTML document
  doc, err := goquery.NewDocumentFromReader(res.Body)
  if err != nil {
		log.Println(err)
		return tickers, err
	}
	
	selector := "#ctl00_contentPlaceHolderConteudo_grdResumoCarteiraTeorica_ctl00 > tbody > tr > td:nth-child(1) > span"

  doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		ticker := s.Text()
		tickers = append(tickers, ticker)
  })
	return tickers, nil
}

// GetAssetData retorna o resultado do scraping dado o ticker do ativo
func (s *WebScrapingService) GetAssetData(ticker string) (*model.FinancialAsset, error) {
	url := "http://br.advfn.com/common/search/exchanges/quote"

	payload := strings.NewReader(fmt.Sprintf("symbol=BOV:%s&symbol_ok=OK", ticker))

	req, err := http.NewRequest("POST", url, payload)
  if err != nil {
		log.Printf("Erro ao criar requisição POST: %s\n", err.Error())
		return nil, err
	}

	req.Header.Add("cookie", "__cfduid=d4486b47d74319c115e613dcd59fdfc8a1596921832; ADVFNUID=d5eb58f2cfa3d0ff2de8ce756e0640b3aaed3d7; recent_stocks=BOV%255EITUB3%252CBOV%255EPETR4")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
  if err != nil {
		log.Printf("Erro ao enviar requisição: %s\n", err.Error())
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatalf("Erro ao realizar parse do documento: %s\n", err.Error())
	}

	asset, err := scrapeAsset(doc)
  if err != nil {
		log.Fatalf("Erro ao scraping do documento: %s\n", err.Error())
		return nil, err
	}

	return asset, nil
}

// GetAssetDataCh canaliza o resultado do scraping dado o ticker do ativo
func (s *WebScrapingService) GetAssetDataCh(ticker string, dataCh chan *model.FinancialAsset, errCh chan error) {
	log.Printf("Começando a busca de dados do ticker %s\n", ticker)
	url := "http://br.advfn.com/common/search/exchanges/quote"

	payload := strings.NewReader(fmt.Sprintf("symbol=BOV:%s&symbol_ok=OK", ticker))

	req, err := http.NewRequest("POST", url, payload)
  if err != nil {
		log.Printf("Erro ao criar requisição POST: %s\n", err.Error())
		errCh <- err
	}
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
  if err != nil {
		log.Printf("Erro ao enviar requisição: %s\n", err.Error())
		errCh <- err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Printf("Erro ao realizar parse do documento: %s\n", err.Error())
		errCh <- err
	}

	asset, err := scrapeAsset(doc)
  if err != nil {
		log.Printf("Erro ao scraping do documento: %s\n", err.Error())
		errCh <- err
	}
	log.Printf("Finalizado a busca de dados do ativo: %s\n", asset.Ticker)
	dataCh <- asset
}

func scrapeAsset(doc *goquery.Document) (*model.FinancialAsset, error) {
	asset := &model.FinancialAsset{}

	doc.Find(tickerSelector).Each(func(i int, s *goquery.Selection) {
		asset.Ticker = strings.TrimSpace(s.Text())
	})

	doc.Find(companyNameSelector).Each(func(i int, s *goquery.Selection) {
		asset.Company = strings.TrimSpace(s.Text())
	})

	doc.Find(currentQuotationSelector).Each(func(i int, s *goquery.Selection) {
		text := strings.Replace(strings.TrimSpace(s.Text()), ",", ".", 1)
		currentPrice, err := strconv.ParseFloat(text, 32)
		if err != nil {
			log.Printf("Erro ao buscar o dado sobre o preço atual: %s\n", err.Error())
		}
		asset.Price = float32(currentPrice)
	})

	doc.Find(openQuotationSelector).Each(func(i int, s *goquery.Selection) {
		text := strings.Replace(strings.TrimSpace(s.Text()), ",", ".", 1)
		openPrice, err := strconv.ParseFloat(text, 32)
		if err != nil {
			log.Printf("Erro ao buscar o dado sobre o preço de abertura: %s\n", err.Error())
		}
		asset.Open = float32(openPrice)
	})

	doc.Find(closeQuotationSelector).Each(func(i int, s *goquery.Selection) {
		text := strings.Replace(strings.TrimSpace(s.Text()), ",", ".", 1)
		closePrice, err := strconv.ParseFloat(text, 32)
		if err != nil {
			log.Printf("Erro ao buscar o dado sobre o preço de fechamento: %s\n", err.Error())
		}
		asset.Close = float32(closePrice)
	})

	doc.Find(pctVariationSelector).Each(func(i int, s *goquery.Selection) {
		pctText := strings.Replace(strings.TrimSpace(s.Text()), ",", ".", 1)
		text := strings.TrimSuffix(pctText, "%")
		pctVariation, err := strconv.ParseFloat(text, 32)
		if err != nil {
			log.Printf("Erro ao buscar o dado sobre a variação percentual: %s\n", err.Error())
		}
		asset.PctVariation = float32(pctVariation)
	})

	doc.Find(priceVariationSelector).Each(func(i int, s *goquery.Selection) {
		text := strings.Replace(strings.TrimSpace(s.Text()), ",", ".", 1)
		priceVariation, err := strconv.ParseFloat(text, 32)
		if err != nil {
			log.Printf("Erro ao buscar o dado sobre a variação de preço: %s\n", err.Error())
		}
		asset.PriceVariation = float32(priceVariation)
	})

	return asset, nil
}