package router

import (
	"blog_api/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/Register", handler.Register)

	return r
}
