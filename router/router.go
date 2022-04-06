package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "pong!",
		})
	})

	userApi := r.Group("/")
	{
		userApi.POST("/user/register")
	}

	// 404 信息返回
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 not found")
	})

	return r
}
