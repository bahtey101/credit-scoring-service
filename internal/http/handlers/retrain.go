package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Hanlder) Retrain(ctx *gin.Context) {
	resp, err := h.scoringservice.Retrain(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrain"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"classification_report": resp.ClassificationReport})
}
