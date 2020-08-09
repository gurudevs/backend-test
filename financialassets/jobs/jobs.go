package jobs

type FinancialAssetsJobs interface {

	InitializeAssetQuotations() (error)

	RefreshAssetQuotations() (error)

}