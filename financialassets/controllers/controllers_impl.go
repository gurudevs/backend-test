package controllers

import (
	"net/http"

	"github.com/ferkze/backend-test/financialassets/usecases"
	"github.com/ferkze/backend-test/utils"
)

// assetsHandler handler do serviço REST de ativos financeiros
type assetsHandler struct {
	assetsUsecases usecases.FinancialAssetsUsecases
}

// NewFinancialAssetsHandler configuração das rotas do handler de ativos financeiros
func NewFinancialAssetsHandler(ucs usecases.FinancialAssetsUsecases) FinancialAssetsController {
	return &assetsHandler{
		assetsUsecases: ucs,
	}
}

// GetAssetsOrderedByVariation handler de busca de ativos financeiros ordenados por maior variação
func (h *assetsHandler) GetAssetsOrderedByVariation(w http.ResponseWriter, r *http.Request) {
	assets, err := h.assetsUsecases.GetAssetsOrderedByVariation()
	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, map[string]string{"message": "Erro interno"})
		return
	}

	utils.JSON(w, http.StatusOK, map[string]interface{}{"data": assets})
	return
}