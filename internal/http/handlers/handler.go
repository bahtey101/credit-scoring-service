package handlers

import (
	"time"

	"github.com/gin-contrib/cors"
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

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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
