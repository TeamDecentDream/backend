package evaluation

import (
	"backend/internal/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func PostEvaluateHandler(c *gin.Context) {
	var input models.EvaluationInput
	err := c.BindJSON(&input)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	err = saveEvaluation(&input)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func GetEvaluationHandler(c *gin.Context) {

}

func DeleteEvaluationHandler(c *gin.Context) {
	var evaluationId int
	err := c.BindJSON(&evaluationId)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"err": "Missing Id"})
		return
	}
}
