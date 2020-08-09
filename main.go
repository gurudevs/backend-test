package main

import (
	"encoding/json"
	"net/http"

	"github.com/ferkze/backend-test/bootstrap"
	"github.com/ferkze/backend-test/financialassets/controllers"
	"github.com/ferkze/backend-test/financialassets/jobs"
	"github.com/ferkze/backend-test/financialassets/repositories/memory"
	"github.com/ferkze/backend-test/financialassets/services/webscraping"
	"github.com/ferkze/backend-test/financialassets/usecases"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	repo := memory.NewFinancialAssetRepository()
	srv := webscraping.NewFinancialAssetScraperService()
	ucs := usecases.NewFinancialAssetsUsecases(srv, repo)
	handlers := controllers.NewFinancialAssetsHandler(ucs)
	tasks := jobs.NewFinancialAssetsJobs(ucs)

	r.HandleFunc("/api/assets-by-variation", handlers.GetAssetsOrderedByVariation)

	bootstrap.SetupCronJobs(tasks)
	bootstrap.SetupApp(ucs)
	bootstrap.SetupRestAPI(r)
}