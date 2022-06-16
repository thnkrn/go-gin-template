package api

import (
	"github.com/gin-gonic/gin"
	handler "github.com/thnkrn/go-gin-template/pkg/api/handler"
)

type Handlers struct {
	Mhandler *handler.MockHandler
}

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(handlers Handlers) *ServerHTTP {
	engine := gin.New()

	// Use logger from Gin
	engine.Use(gin.Logger())

	// Healthcheck
	engine.GET("healthcheck", func(c *gin.Context) {
		c.String(200, "OK")
	})

	api := engine.Group("/api")
	api.GET("mock", handlers.Mhandler.Mock)

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
