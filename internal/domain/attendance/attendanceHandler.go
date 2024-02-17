package attendance

import (
	"backend/internal/utils/jwt"
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func EnterHandler(c *gin.Context) {
	token := c.GetHeader("Authorization")
	id, _, _, _, _ := jwt.AccessTokenVerifier(token)
	timeLog, err := enter(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"state": timeLog.Format("2006-01-02 15:04:05"),
	})
}

func LeaveHandler(c *gin.Context) {
	token := c.GetHeader("Authorization")
	id, _, _, _, _ := jwt.AccessTokenVerifier(token)
	err := leave(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func GetWorkStateHandler(c *gin.Context) {
	token := c.GetHeader("Authorization")
	id, _, _, _, _ := jwt.AccessTokenVerifier(token)
	timeLog, err := GetWorkState(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, gin.H{})
			return
		}

		log.Println("Error state:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"state": timeLog.Format("2006-01-02 15:04:05"),
	})

}
