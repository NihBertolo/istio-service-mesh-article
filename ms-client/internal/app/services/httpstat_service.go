package services

import (
	"ms-client-istioarticle/internal/infrastructure/clients"

	"github.com/gin-gonic/gin"
)

type HTTPStatsService struct {
	client *clients.HTTPStatsClient
}

func NewHTTPStatsService(client *clients.HTTPStatsClient) *HTTPStatsService {
	return &HTTPStatsService{client: client}
}

func (s *HTTPStatsService) GetStats(ctx *gin.Context, param1, param2 int) {
	s.client.GetStats(ctx, param1, param2)
}
