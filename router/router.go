package router

import (
	v1 "MedicalCare/api/v1"
	"MedicalCare/docs"
	"MedicalCare/middleware"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "pong!",
		})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// no token router
	apiUser := r.Group("/api/v1/user")
	{
		apiUser.POST("/register", v1.UserRegister)
		apiUser.POST("/login", v1.UserLogin)
	}

	// token router
	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT())
	{
		// User service
		apiv1.GET("/user/token", v1.RefreshAccessToken)
		apiv1.GET("/user/info", v1.GetUserInfo)
		apiv1.GET("/user/online", v1.UserOnline)
		apiv1.PUT("/user/info", v1.UpdateUserInfo)
		apiv1.POST("/user/password", v1.ResetUserPassword)

		// blog service
		apiv1.POST("/blog", v1.WriteArticle)
		apiv1.GET("/blog", v1.GetArticles)
		apiv1.GET("/blog/:bid", v1.GetArticle)
		apiv1.PUT("/blog/:bid", v1.UpdateArticle)
		apiv1.DELETE("/blog/:bid", v1.RemoveArticle)

		// chat group service
		apiv1.POST("/group", v1.CreateGroup)
		apiv1.PUT("/group/member", v1.InviteMember)

		// chat room service
		apiv1.GET("/chat/:receiver", v1.Chat)
	}

	// 404 Message Return
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 not found")
	})

	return r
}
