package main

import (
	"encoding/json"
	"net/http"

	"github.com/ferkze/backend-test/financialassets/controllers"
	"github.com/ferkze/backend-test/financialassets/repositories/memory"
	"github.com/ferkze/backend-test/financialassets/usecases"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	repo := memory.NewFinancialAssetRepository()
	ucs := usecases.NewFinancialAssetsUsecases(repo)
	handler := controllers.NewFinancialAssetsHandler(ucs)

	r.HandleFunc("/assets-by-variation", handler.GetAssetsOrderedByVariation)

	http.ListenAndServe(":8080", r)
}