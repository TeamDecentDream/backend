package significant

import (
	"backend/internal/models"
	"backend/internal/utils/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSignificantCountHandler(c *gin.Context) {

}

func GetSignificantsHandler(c *gin.Context) {

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

func PatchSignificantHandler(c *gin.Context) {

}

func DeleteSignificantHandler(c *gin.Context) {

}
