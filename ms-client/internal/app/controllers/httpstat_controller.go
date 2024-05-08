package controllers

import (
	"ms-client-istioarticle/internal/app/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HTTPStatsController struct {
	statsService *services.HTTPStatsService
}

func NewHTTPStatsController(statsService *services.HTTPStatsService) *HTTPStatsController {
	return &HTTPStatsController{statsService: statsService}
}

// GetStatsHandler godoc
// @Summary Proxy HTTPSTATS
// @Description Get httpstats response
// @Tags httpstats
// @Accept json
// @Produce json
// @Param status query string false "Status de Resposta" Format(int32)
// @Param delay query int false "Delay para resposta da API" Format(delay)
// @Success 200 {object} dto.JSONResponse
// @Router /httpstats [get]
func (c *HTTPStatsController) GetStatsHandler(ctx *gin.Context) {
	param1Str := ctx.Query("status")
	status, err := strconv.Atoi(param1Str)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "status deve ser um número inteiro"})
		return
	}

	param2Str := ctx.Query("delay")
	delay, err := strconv.Atoi(param2Str)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "delay deve ser um número inteiro"})
		return
	}

	c.statsService.GetStats(ctx, status, delay)
}
