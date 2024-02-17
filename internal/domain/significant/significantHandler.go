package significant

import (
	"backend/internal/models"
	"backend/internal/utils/jwt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

func GetSignificantCountHandler(c *gin.Context) {
	count, err := GetSignificantCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": count})
}

func GetSignificantHandler(c *gin.Context) {
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
	siginificants, err := getSignificants(_page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusOK, siginificants)
}

func PostSignificantHandler(c *gin.Context) {
	token := c.GetHeader("Authorization")
	id, _, _, _, err := jwt.AccessTokenVerifier(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization Code Error"})
		return
	}
	var inputSignificant models.Significant
	err = c.BindJSON(&inputSignificant)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input Value Error"})
		return
	}
	inputSignificant.AuthorID = id
	err = saveSignificant(&inputSignificant)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func PutSignificantHandler(c *gin.Context) {
	token := c.GetHeader("Authorization")
	id, _, _, authorities, err := jwt.AccessTokenVerifier(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization Code Error"})
		return
	}

	var inputSignificant models.SignificantOutput
	err = c.BindJSON(&inputSignificant)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input Value Error"})
		return
	}

	existingSignificant, err := findSignificantById(inputSignificant.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	log.Println(existingSignificant)
	flag := false
Loop1:
	for _, authority := range authorities {
		if "ROLE_ADMIN" == authority.Role {
			flag = true
			break Loop1
		}
	}

	if id == existingSignificant.AuthorID || flag {
		loc, err := time.LoadLocation("Asia/Seoul")
		updateTime := time.Now().In(loc)

		significant := models.Significant{
			ID:         inputSignificant.ID,
			Contents:   inputSignificant.Contents,
			Warn:       inputSignificant.Warn,
			AuthorID:   id,
			RegDate:    inputSignificant.RegDate,
			UpdateDate: updateTime,
		}
		err = updateSignificant(&significant)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"error": err})
}

func DeleteSignificantHandler(c *gin.Context) {
	token := c.GetHeader("Authorization")
	_id, flag := c.GetQuery("id")
	if !flag {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad Request"})
		return
	}
	_SId, err := strconv.Atoi(_id)
	id, _, _, authorities, err := jwt.AccessTokenVerifier(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization Code Error"})
		return
	}

	for _, authority := range authorities {
		if "ROLE_ADMIN" == authority.Role {
			_ = DeleteSignificantByAdmin(_SId)
			c.JSON(http.StatusOK, gin.H{})
			return
		}
	}
	_ = DeleteSignificant(_SId, id)
	c.JSON(http.StatusOK, gin.H{})
}
