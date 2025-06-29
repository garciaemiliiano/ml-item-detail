package ping

import (
	"item-detail-api/src/entrypoints/rest"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PingHandler struct {
	rest.HandlerBase
	logger *zap.Logger
}

func NewWirePingHandler(logger *zap.Logger) PingHandler {
	return PingHandler{rest.HandlerBase{}, logger}
}

func NewPingHandler(logger *zap.Logger) PingHandler {
	return PingHandler{rest.HandlerBase{}, logger}
}

func (h PingHandler) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, rest.Okf("pong"))
	}
}
