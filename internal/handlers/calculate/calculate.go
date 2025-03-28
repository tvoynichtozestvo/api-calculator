package calculate

import (
	"api-calculator/internal/models"
	"api-calculator/internal/services/logic"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CalculateHandler(c *gin.Context) {
	var req models.CalculateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := logic.Calculate(req.Operation, req.A, req.B)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.CalculateResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.CalculateResponse{Result: result})
}
