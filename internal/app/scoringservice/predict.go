package scoringservice

import (
	"context"

	"github.com/bahtey101/credit-scoring-service/internal/model"
)

func (s *ScoringService) Predict(
	ctx context.Context,
	features model.Features,
) (model.Prediction, error) {
	var prediction model.Prediction

	if err := s.modelClient.Post(nil, "/predict", features, &prediction); err != nil {
		return model.Prediction{}, err
	}

	return prediction, nil
}
