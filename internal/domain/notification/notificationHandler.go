package notification

import (
	"backend/internal/models"
	"backend/internal/utils/jwt"
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
	token := c.GetHeader("Authorization")
	id, _, _, _, err := jwt.AccessTokenVerifier(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization Code Error"})
		return
	}
	var inputNotification models.Notification
	err = c.BindJSON(&inputNotification)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input Value Error"})
		return
	}
	inputNotification.AuthorID = id
	err = saveNotification(&inputNotification)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})

}

func PutNotificationHandler(c *gin.Context) {
	token := c.GetHeader("Authorization")
	id, _, _, authorities, err := jwt.AccessTokenVerifier(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization Code Error"})
		return
	}

	var inputNotification models.Notification
	err = c.BindJSON(&inputNotification)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input Value Error"})
		return
	}

	existingNotification, err := findNotificationById(inputNotification.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	flag := false
Loop1:
	for _, authority := range authorities {
		if "ROLE_ADMIN" == authority.Role {
			flag = true
			break Loop1
		}
	}
	if existingNotification.AuthorID == id || flag {
		err = updateNotification(&inputNotification, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"error": err})
}

func DeleteNotificationHandler(c *gin.Context) {
	token := c.GetHeader("Authorization")
	_id, flag := c.GetQuery("id")
	if !flag {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Missing PostId"})
		return
	}
	id, err := strconv.Atoi(_id)

	Authorid, _, _, authorities, err := jwt.AccessTokenVerifier(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization Code Error"})
		return
	}

	for _, authority := range authorities {
		if "ROLE_ADMIN" == authority.Role {
			_ = DeleteNotificationByAdmin(id)
			c.JSON(http.StatusOK, gin.H{})
			return
		}
	}
	_ = DeleteNotification(id, Authorid)
	c.JSON(http.StatusOK, gin.H{})
}
