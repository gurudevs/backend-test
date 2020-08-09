package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/ferkze/backend-test/financialassets/controllers"
	"github.com/ferkze/backend-test/financialassets/repositories/memory"
	"github.com/ferkze/backend-test/financialassets/services/webscraping"
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
	srv := webscraping.NewFinancialAssetScraperService()
	ucs := usecases.NewFinancialAssetsUsecases(srv, repo)
	handlers := controllers.NewFinancialAssetsHandler(ucs)

	r.HandleFunc("/api/assets-by-variation", handlers.GetAssetsOrderedByVariation)


	port := os.Getenv("port")
	fmt.Printf("Server listening on port %s...\n", port)
	http.ListenAndServe(":"+port, r)
}