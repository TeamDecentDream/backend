package config

import (
	"backend/internal/domain/attendance"
	"backend/internal/domain/evaluation"
	"backend/internal/domain/member"
	"backend/internal/domain/notification"
	"backend/internal/domain/significant"
	"backend/internal/domain/todo"
	"backend/internal/domain/transaction"
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
	setAttendanceRoute(router)
	setEvaluationRoute(router)
	setTransactionRoute(router)
	setTodoRoute(router)
	return router
}

func setMemberRoute(router *gin.Engine) {
	router.GET("/member", middleware.AuthMiddleware([]string{"ROLE_ADMIN", "ROLE_PRIME"}), member.GetAllMembersHandler)
	router.GET("/member/details", middleware.AuthMiddleware([]string{"ROLE_ADMIN", "ROLE_WORKER", "ROLE_GUEST", "ROLE_PRIME"}), member.GetMemberDetailHandler)
	router.GET("/member/count", middleware.AuthMiddleware([]string{"ROLE_ADMIN", "ROLE_PRIME"}), member.GetMemberCountHandler)
	router.POST("/member/login", member.LoginHandler)
	router.PUT("/member/confirm", middleware.AuthMiddleware([]string{"ROLE_ADMIN"}), member.ConfirmHandler)
	router.DELETE("/member", member.DeleteMemberHandler)
	router.POST("/member/address", member.ConnectAddressHandler)
}

func setNotificationRoute(router *gin.Engine) {
	router.GET("/notification", middleware.AuthMiddleware([]string{"ROLE_ADMIN", "ROLE_WORKER", "ROLE_GUEST", "ROLE_PRIME"}), notification.GetNotificationHandler)
	router.GET("/notification/count", middleware.AuthMiddleware([]string{"ROLE_ADMIN", "ROLE_WORKER", "ROLE_GUEST", "ROLE_PRIME"}), notification.GetNotificationCountHandler)
	router.POST("/notification", middleware.AuthMiddleware([]string{"ROLE_ADMIN", "ROLE_PRIME"}), notification.PostNotificationHandler)
	router.PUT("/notification", middleware.AuthMiddleware([]string{"ROLE_ADMIN", "ROLE_PRIME"}), notification.PutNotificationHandler)
	router.DELETE("/notification", middleware.AuthMiddleware([]string{"ROLE_ADMIN", "ROLE_PRIME"}), notification.DeleteNotificationHandler)
}

func setSignificantRoute(router *gin.Engine) {
	router.GET("/significant", middleware.AuthMiddleware([]string{"ROLE_ADMIN", "ROLE_WORKER", "ROLE_PRIME"}), significant.GetSignificantHandler)
	router.GET("/significant/count", middleware.AuthMiddleware([]string{"ROLE_ADMIN", "ROLE_WORKER", "ROLE_PRIME"}), significant.GetSignificantCountHandler)
	router.POST("/significant", middleware.AuthMiddleware([]string{"ROLE_ADMIN", "ROLE_WORKER", "ROLE_PRIME"}), significant.PostSignificantHandler)
	router.PUT("/significant", middleware.AuthMiddleware([]string{"ROLE_ADMIN", "ROLE_WORKER", "ROLE_PRIME"}), significant.PutSignificantHandler)
	router.DELETE("/significant", middleware.AuthMiddleware([]string{"ROLE_ADMIN", "ROLE_WORKER", "ROLE_PRIME"}), significant.DeleteSignificantHandler)
}

func setAttendanceRoute(router *gin.Engine) {
	router.POST("/attendance/enter", middleware.AuthMiddleware([]string{"ROLE_ADMIN", "ROLE_WORKER", "ROLE_PRIME"}), attendance.EnterHandler)
	router.POST("/attendance/leave", middleware.AuthMiddleware([]string{"ROLE_ADMIN", "ROLE_WORKER", "ROLE_PRIME"}), attendance.LeaveHandler)
	router.GET("/attendance", middleware.AuthMiddleware([]string{"ROLE_ADMIN", "ROLE_WORKER", "ROLE_PRIME"}), attendance.GetWorkStateHandler)
	router.GET("/attendance/timelogs", middleware.AuthMiddleware([]string{"ROLE_ADMIN", "ROLE_WORKER", "ROLE_PRIME"}), attendance.GetWorkLogHandler)
}

func setEvaluationRoute(router *gin.Engine) {
	router.POST("/evaluate", middleware.AuthMiddleware([]string{"ROLE_PRIME"}), evaluation.PostEvaluateHandler)
	router.GET("/evaluate", middleware.AuthMiddleware([]string{"ROLE_PRIME", "ROLE_ADMIN"}), evaluation.GetEvaluationHandler)
	router.DELETE("/evaluate", middleware.AuthMiddleware([]string{"ROLE_PRIME", "ROLE_ADMIN"}), evaluation.DeleteEvaluationHandler)
}

func setTransactionRoute(router *gin.Engine) {
	router.POST("/transaction", middleware.AuthMiddleware([]string{"ROLE_ADMIN"}), transaction.PostTransactionHandler)
	router.GET("/transaction", middleware.AuthMiddleware([]string{"ROLE_ADMIN", "ROLE_WORKER", "ROLE_PRIME"}), transaction.GetTransactionHandler)
	router.GET("/transaction/count", middleware.AuthMiddleware([]string{"ROLE_ADMIN", "ROLE_WORKER", "ROLE_PRIME"}), transaction.GetTransactionCountHandler)
	router.PUT("/transaction", middleware.AuthMiddleware([]string{"ROLE_ADMIN"}), transaction.PutTransactionHandler)
	router.DELETE("/transaction", middleware.AuthMiddleware([]string{"ROLE_ADMIN"}), transaction.DeleteTransactionHandler)
}

func setTodoRoute(router *gin.Engine) {
	router.GET("/todo", middleware.AuthMiddleware([]string{"ROLE_ADMIN", "ROLE_WORKER", "ROLE_PRIME"}), todo.GetTodoHandler)
	router.POST("/todo", middleware.AuthMiddleware([]string{"ROLE_ADMIN", "ROLE_WORKER", "ROLE_PRIME"}), todo.PostTodoHandler)
	router.PUT("/todo", middleware.AuthMiddleware([]string{"ROLE_ADMIN", "ROLE_WORKER", "ROLE_PRIME"}), todo.PUTTodoHandler)
	router.DELETE("/todo", middleware.AuthMiddleware([]string{"ROLE_ADMIN", "ROLE_WORKER", "ROLE_PRIME"}), todo.DeleteTodoHandler)
}
