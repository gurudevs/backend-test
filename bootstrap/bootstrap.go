package bootstrap

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ferkze/backend-test/financialassets/jobs"
	"github.com/gorilla/mux"
	cron "github.com/robfig/cron/v3"
)

// SetupRestAPI Configura a execução do servidor de API RESTFUL
func SetupRestAPI(r *mux.Router) {
	port := os.Getenv("port")
	fmt.Printf("Server listening on port %s...\n", port)
	http.ListenAndServe(":"+port, r)
}

// SetupCronJobs Configura a execução das tarefas do servidor
func SetupCronJobs(tasks jobs.FinancialAssetsJobs) {
	c := cron.New()
	// Inicializa cotações do dia às 8h30 de segunda a sexta
	c.AddFunc("0 30 8 * * 1-5", tasks.InitializeAssetQuotations())
	// Atualiza cotações a cada 15 minutos de segunda a sexta
	c.AddFunc("0 */15 9-17 * * 1-5", tasks.RefreshAssetQuotations())
	
}