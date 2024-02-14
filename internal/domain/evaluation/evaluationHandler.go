package evaluation

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func EvaluateHandler(c *gin.Context) {
	var input map[string]interface{}
	err := c.BindJSON(&input)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	for key, value := range input {
		log.Printf("%s: %v\n", key, value)
	}
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
