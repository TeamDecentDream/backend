package notification

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetNotificationCountHandler(c *gin.Context) {
	count, err := GetNotificationCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"count": count})
}

func GetNotificationHandler(c *gin.Context) {
	page := c.Query("page")
	if page == "" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	_page, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	notifications, err := getNotification(_page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusOK, notifications)

}

func PostNotificationHandler(c *gin.Context) {

}

func PatchNotificationHandler(c *gin.Context) {

}

func DeleteNotificationHandler(c *gin.Context) {

}
