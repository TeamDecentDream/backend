package config

import (
	"backend/internal/domain/member"
	"backend/internal/domain/notification"
	"backend/internal/domain/significant"
	"backend/internal/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
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
	router.GET("/member", middleware.AuthMiddleware([]string{"ROLE_ADMIN"}), member.GetAllMembersHandler)
	router.GET("/member/details", middleware.AuthMiddleware([]string{"ROLE_ADMIN", "ROLE_WORKER", "ROLE_GUEST", "ROLE_PRIME"}), member.GetMemberDetailHandler)
	router.POST("/member/login", member.LoginHandler)
	router.DELETE("/member", member.DeleteMemberHandler)
	router.POST("/address", member.ConnectAddressHandler)
}

func setNotificationRoute(router *gin.Engine) {
	router.GET("/notification", middleware.AuthMiddleware([]string{"ROLE_ADMIN", "ROLE_WORKER", "ROLE_GUEST", "ROLE_PRIME"}), notification.GetNotificationHandler)
	router.GET("/notification/count", middleware.AuthMiddleware([]string{"ROLE_ADMIN", "ROLE_WORKER", "ROLE_GUEST", "ROLE_PRIME"}), notification.GetNotificationCountHandler)
	router.POST("/notification", middleware.AuthMiddleware([]string{"ROLE_ADMIN", "ROLE_PRIME"}), notification.PostNotificationHandler)
	router.PATCH("/notification", middleware.AuthMiddleware([]string{"ROLE_ADMIN", "ROLE_PRIME"}), notification.PatchNotificationHandler)
	router.DELETE("/notification", middleware.AuthMiddleware([]string{"ROLE_ADMIN", "ROLE_PRIME"}), notification.DeleteNotificationHandler)
}

func setSignificantRoute(router *gin.Engine) {
	router.GET("/significant", middleware.AuthMiddleware([]string{"ROLE_ADMIN", "ROLE_WORKER", "ROLE_GUEST", "ROLE_PRIME"}), significant.GetSignificantHandler)
	router.GET("/significant/count", middleware.AuthMiddleware([]string{"ROLE_ADMIN", "ROLE_WORKER", "ROLE_GUEST", "ROLE_PRIME"}), significant.GetSignificantCountHandler)
	router.POST("/significant", middleware.AuthMiddleware([]string{"ROLE_ADMIN", "ROLE_WORKER", "ROLE_GUEST", "ROLE_PRIME"}), significant.PostSignificantHandler)
	router.PATCH("/significant", middleware.AuthMiddleware([]string{"ROLE_ADMIN", "ROLE_WORKER", "ROLE_GUEST", "ROLE_PRIME"}), significant.PatchSignificantHandler)
	router.DELETE("/significant", middleware.AuthMiddleware([]string{"ROLE_ADMIN", "ROLE_WORKER", "ROLE_GUEST", "ROLE_PRIME"}), significant.DeleteSignificantHandler)
}
