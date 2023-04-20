package server

import (
	"github.com/gin-gonic/gin"
	"github.com/lbernardo/fx-webserver/internal/providers/server/handlers"
)

func RegisterRoutes(router *gin.RouterGroup, handler *handlers.Handler) {
	router.GET("/", handler.Get)
	router.POST("/", handler.Create)
	router.DELETE("/:id", handler.Delete)
	router.PATCH("/:id", handler.Update)
}
