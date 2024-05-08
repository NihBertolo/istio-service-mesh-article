package clients

import (
	"fmt"
	"io/ioutil"
	"ms-client-istioarticle/internal/app/dto"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HTTPStatsClient struct {
	baseURL string
}

func NewHTTPStatsClient(baseURL string) *HTTPStatsClient {
	return &HTTPStatsClient{baseURL: baseURL}
}

func (c *HTTPStatsClient) GetStats(ctx *gin.Context, status, delay int) {
	url := fmt.Sprintf("%s/%d?sleep=%d", c.baseURL, status, delay)
	startTime := time.Now()

	response, err := http.Get(url)
	if err != nil {
		dto.FailResponse(ctx, 501, err.Error(), 0)
	}
	defer response.Body.Close()

	endTime := time.Now()
	responseTime := endTime.Sub(startTime)

	bodyMessage, err := ioutil.ReadAll(response.Body)
	if err != nil {
		dto.FailResponse(ctx, 500, err.Error(), responseTime.Seconds())
	}

	if response.StatusCode >= 400 {
		dto.FailResponse(ctx, response.StatusCode, string(bodyMessage), responseTime.Seconds())
	} else {
		dto.SuccessResponse(ctx, responseTime.Seconds(), response.StatusCode)
	}
}
