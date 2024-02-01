package config

import (
	"backend/internal/domain/member"
	"backend/internal/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://gryu-dev.com", "http://gryu-dev.com"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
	}))

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"msg": "pong"})
	})
	setMemberRoute(router)
	setNotificationRoute(router)
	setSignificantRoute(router)
	return router
}

func setMemberRoute(router *gin.Engine) {
	router.GET("/member", middleware.AuthMiddleware([]string{"ROLE_ADMIN"}), member.GetAllMembers)
	router.GET("/member/details", middleware.AuthMiddleware([]string{"ROLE_ADMIN", "ROLE_WORKER", "ROLE_GUEST", "ROLE_PRIME"}), member.GetMemberDetail)
	router.POST("/member/login", member.Login)
	router.PATCH("/member/details", member.UpdateMember)
	router.DELETE("/member", member.DeleteMember)
	router.POST("/address", member.ConnectAddress)
}

func setNotificationRoute(router *gin.Engine) {

}

func setSignificantRoute(router *gin.Engine) {

}
