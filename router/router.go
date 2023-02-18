package router

import (
	"github.com/gin-gonic/gin"
	"tiktok/handler"
)

func InitDouyinRouter() *gin.Engine {
	r := gin.Default()

	r.Static("static", "./static")

	baseGroup := r.Group("/douyin")

	baseGroup.POST("/user/register/", handler.RegisterHandler)

	return r
}
