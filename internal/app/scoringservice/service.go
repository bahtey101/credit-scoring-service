package scoringservice

import "github.com/bahtey101/credit-scoring-service/internal/config"

type ScoringService struct {
	modelHost string
	modelPort string
}

func NewScoringService(cfg *config.Config) *ScoringService {
	return &ScoringService{
		modelHost: cfg.MLHost,
		modelPort: cfg.MLPort,
	}
}
