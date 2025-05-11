package model

type Features struct {
	Utilization     float32 `json:"RevolvingUtilizationOfUnsecuredLines"`
	Age             float32 `json:"age"`
	PastDueNotWorse float32 `json:"NumberOfTime30_59DaysPastDueNotWorse"`
	DebtRatio       float32 `json:"DebtRatio"`
	MonthlyIncome   float32 `json:"MonthlyIncome"`
	OpenLines       float32 `json:"NumberOfOpenCreditLinesAndLoans"`
	RealLines       float32 `json:"NumberRealEstateLoansOrLines"`
	Dependents      float32 `json:"NumberOfDependents"`
}
