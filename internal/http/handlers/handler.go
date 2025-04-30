package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/bahtey101/credit-scoring-service/internal/app/scoringservice"
)

type Hanlder struct {
	secret string

	scoringservice *scoringservice.ScoringService
}

func NewHandler(
	secret string,
	scoringservice *scoringservice.ScoringService,
) *Hanlder {
	return &Hanlder{
		secret:         secret,
		scoringservice: scoringservice,
	}
}

func (h *Hanlder) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group(("/api"))
	{
		scoringGroup := api.Group("/scoring")
		{
			scoringGroup.POST("/predict", h.Predict)
			scoringGroup.POST("/retrain", h.Retrain)
		}
	}

	return router
}
