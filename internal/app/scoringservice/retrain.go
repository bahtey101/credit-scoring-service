package scoringservice

import (
	"github.com/bahtey101/credit-scoring-service/internal/model"
	"github.com/gin-gonic/gin"
)

func (s *ScoringService) Retrain(ctx *gin.Context) (model.RetrainResponse, error) {
	var report model.RetrainResponse
	if err := s.modelClient.Post(nil, "/retrain", nil, &report); err != nil {
		return model.RetrainResponse{}, err
	}
	return report, nil
}
