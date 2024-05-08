package dto

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JSONResponse struct {
	Code      int     `json:"code" example:"200"`
	Message   string  `json:"message" example:"Success"`
	Timestamp float64 `json:"timestamp"`
}

func SuccessResponse(c *gin.Context, timestamp float64, respCode int) {
	c.JSON(respCode, JSONResponse{
		Code:      respCode,
		Timestamp: timestamp,
		Message:   "Success",
	})
}

func FailResponse(c *gin.Context, respCode int, message string, timestamp float64) {
	if respCode == http.StatusInternalServerError {
		c.JSON(respCode, JSONResponse{
			Code:      respCode,
			Timestamp: timestamp,
			Message:   message,
		})
		return
	}

	c.JSON(respCode, JSONResponse{
		Code:      respCode,
		Timestamp: timestamp,
		Message:   message,
	})
}
