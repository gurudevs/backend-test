package bootstrap

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ferkze/backend-test/financialassets/jobs"
	"github.com/ferkze/backend-test/financialassets/usecases"
	"github.com/gorilla/mux"
	cron "github.com/robfig/cron/v3"
)

// SetupApp Inicializa a aplicação com o caso de uso inicial para carregar os dados
func SetupApp(ucs usecases.FinancialAssetsUsecases) {
	err := ucs.InitializeAssetQuotations()
	if err != nil {
		log.Fatalf("A aplicação falhou em sua inicialização: %s\n", err.Error())
		return
	}
	log.Println("A aplicação inicializou com sucesso!")
}

// SetupRestAPI Configura a execução do servidor de API RESTFUL
func SetupRestAPI(r *mux.Router) {
	port := os.Getenv("port")
	fmt.Printf("O servidor está no ar, acesse http://localhost:%s\n", port)
	http.ListenAndServe(":"+port, r)
}

// SetupCronJobs Configura a execução das tarefas do servidor
func SetupCronJobs(tasks jobs.FinancialAssetsJobs) {
	c := cron.New()
	
	// Inicializa cotações do dia às 8h30 de segunda a sexta
	c.AddFunc("0 30 8 * * 1-5", func() {
		err := tasks.InitializeAssetQuotations()
		if err != nil {
			log.Printf("Erro na execução do Job tasks.InitializeAssetQuotations(): %s\n", err.Error())
			return
		}
		log.Println("Tarefa tasks.InitializeAssetQuotations() executada com sucesso!")
	})

	// Atualiza cotações a cada 15 minutos de segunda a sexta
	c.AddFunc("0 */15 9-17 * * 1-5", func() {
		err := tasks.RefreshAssetQuotations()
		if err != nil {
			log.Printf("Erro na execução do Job tasks.RefreshAssetQuotations(): %s\n", err.Error())
			return
		}
		log.Println("Tarefa tasks.RefreshAssetQuotations() executada com sucesso!")
	})
	
}