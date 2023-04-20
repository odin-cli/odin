package server

import "github.com/gin-gonic/gin"

func NewGinEngine() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	return gin.Default()
}

func NewRouter(engine *gin.Engine) *gin.RouterGroup {
	return engine.Group("/")
}
