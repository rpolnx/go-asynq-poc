package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rpolnx/go-asynq-poc/internal/configs"
)

type HttpServer struct {
	Engine *gin.Engine
}

func InitializeServer(config *configs.AppConfig) (*HttpServer, error) {
	r := gin.New()

	r.Use(gin.Recovery())

	// r.Use(nil){}

	r.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})

	return &HttpServer{
		Engine: r,
	}, nil
}
