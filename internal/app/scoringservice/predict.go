package scoringservice

import (
	"context"

	"github.com/bahtey101/credit-scoring-service/internal/model"
)

func (s *ScoringService) Predict(
	ctx context.Context,
	features model.Features,
) (model.Prediction, error) {

	return model.Prediction{
		Value: 0.5,
	}, nil
}
