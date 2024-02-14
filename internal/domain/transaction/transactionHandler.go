package transaction

import (
	"backend/internal/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func DeleteTransactionHandler(c *gin.Context) {
	var transactionId int
	err := c.BindJSON(&transactionId)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"err": "Missing PostId"})
		return
	}
	err = DeleteTransaction(transactionId)
	if err != nil {
		log.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func PutTransactionHandler(c *gin.Context) {
	var transaction models.Transaction
	err := c.BindJSON(&transaction)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"err": "Missing Params"})
		return
	}
	err = PutTransaction(&transaction)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func PostTransactionHandler(c *gin.Context) {
	var transaction models.Transaction
	err := c.BindJSON(&transaction)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"err": "Missing Params"})
		return
	}
	err = PostTransaction(&transaction)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func GetTransactionHandler(c *gin.Context) {
	var page int
	_page, flag := c.GetQuery("page")
	if !flag {
		page = 1
	} else {
		page, _ = strconv.Atoi(_page)
	}

	transactions, err := GetTransactions(page)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusOK, gin.H{"transactions": transactions})
}

func GetTransactionCountHandler(c *gin.Context) {
	count, err := GetTransactionCount()
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Sever Error"})
	}
	c.JSON(http.StatusOK, gin.H{"count": count})
}
