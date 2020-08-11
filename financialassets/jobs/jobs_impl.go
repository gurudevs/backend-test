package jobs

import "github.com/ferkze/backend-test/financialassets/usecases"

// assetsJobs Tarefas do serviço de ativos financeiros
type assetsJobsImpl struct {
	assetsUsecases usecases.FinancialAssetsUsecases
}

// NewFinancialAssetsJobs configuração das rotas do Jobs de ativos financeiros
func NewFinancialAssetsJobs(ucs usecases.FinancialAssetsUsecases) FinancialAssetsJobs {
	return &assetsJobsImpl{
		assetsUsecases: ucs,
	}
}

func (j *assetsJobsImpl) InitializeAssetQuotations() (error) {
	return j.assetsUsecases.PopulateAssets()
}

func (j *assetsJobsImpl) RefreshAssetQuotations() (error) {
	return j.assetsUsecases.RefreshAssetQuotations()
}