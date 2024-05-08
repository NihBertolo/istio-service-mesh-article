package main

import (
	_ "ms-client-istioarticle/cmd/docs"
	"ms-client-istioarticle/internal/app/controllers"
	"ms-client-istioarticle/internal/app/services"
	"ms-client-istioarticle/internal/infrastructure/clients"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Resiliência e Tolerância a Falhas com Istio
// @version 1.0
// @description Aplicação desenvolvida para testes do circuit breaker.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api/v1
// @schemes http
func main() {
	client := clients.NewHTTPStatsClient("http://httpstat.us")
	service := services.NewHTTPStatsService(client)

	controller := controllers.NewHTTPStatsController(service)

	router := gin.Default()
	httpstatAPI := router.Group("/api/v1")
	{
		httpstatAPI.GET("/httpstats", controller.GetStatsHandler)
		httpstatAPI.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	router.Run(":8080")
}
