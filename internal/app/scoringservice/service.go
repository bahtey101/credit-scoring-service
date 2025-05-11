package scoringservice

import (
	"fmt"
	"time"

	"github.com/bahtey101/credit-scoring-service/internal/config"
	"github.com/bahtey101/credit-scoring-service/internal/http/client"
)

type ScoringService struct {
	modelClient *client.Client
}

func NewScoringService(
	cfg *config.Config,
	timeout time.Duration,
	retries int,
) *ScoringService {
	return &ScoringService{
		modelClient: client.NewClient(
			fmt.Sprintf("http://%s:%s", cfg.MLHost, cfg.MLPort),
			timeout,
			retries,
		),
	}
}
