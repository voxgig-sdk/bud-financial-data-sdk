package core

var UtilityRegistrar func(u *Utility)

var NewBaseFeatureFunc func() Feature

var NewDebugFeatureFunc func() Feature

var NewIdempotencyFeatureFunc func() Feature

var NewMetricsFeatureFunc func() Feature

var NewPagingFeatureFunc func() Feature

var NewRatelimitFeatureFunc func() Feature

var NewRetryFeatureFunc func() Feature

var NewTestFeatureFunc func() Feature

var NewTimeoutFeatureFunc func() Feature

var NewCorrectFinancialDataEntityFunc func(client *BudFinancialDataSDK, entopts map[string]any) BudFinancialDataEntity

var NewCustomerMerchantCorrectionEntityFunc func(client *BudFinancialDataSDK, entopts map[string]any) BudFinancialDataEntity

var NewLabelEntityFunc func(client *BudFinancialDataSDK, entopts map[string]any) BudFinancialDataEntity

var NewListLabelEntityFunc func(client *BudFinancialDataSDK, entopts map[string]any) BudFinancialDataEntity

var NewListTransactionEntityFunc func(client *BudFinancialDataSDK, entopts map[string]any) BudFinancialDataEntity

var NewManageFinancialDataEntityFunc func(client *BudFinancialDataSDK, entopts map[string]any) BudFinancialDataEntity

var NewManageTransactionLabelEntityFunc func(client *BudFinancialDataSDK, entopts map[string]any) BudFinancialDataEntity

var NewRetrieveFinancialDataEntityFunc func(client *BudFinancialDataSDK, entopts map[string]any) BudFinancialDataEntity

var NewSimilarEntityFunc func(client *BudFinancialDataSDK, entopts map[string]any) BudFinancialDataEntity

