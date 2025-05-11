package handlers

import (
	"net/http"

	"github.com/bahtey101/credit-scoring-service/internal/model"
	"github.com/gin-gonic/gin"
)

func (h *Hanlder) Predict(ctx *gin.Context) {
	var features model.Features
	if err := ctx.ShouldBindJSON(&features); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	prediction, err := h.scoringservice.Predict(ctx, features)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to predict"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"prediction": prediction.Value,
	})
}
