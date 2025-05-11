package model

type ClassificationMetrics struct {
	Precision float64 `json:"precision"`
	Recall    float64 `json:"recall"`
	F1Score   float64 `json:"f1-score"`
	Support   float64 `json:"support"`
}

type RetrainResponse struct {
	ClassificationReport struct {
		NonDefault  ClassificationMetrics `json:"0"`
		Default     ClassificationMetrics `json:"1"`
		Accuracy    float64               `json:"accuracy"`
		MacroAvg    ClassificationMetrics `json:"macro avg"`
		WeightedAvg ClassificationMetrics `json:"weighted avg"`
	} `json:"classification_report"`
}
