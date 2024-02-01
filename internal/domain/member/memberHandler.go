package member

import (
	"backend/internal/db"
	"github.com/gin-gonic/gin"
	"log"
)

func GetAllMembers(c *gin.Context) {
	memberList, err := db.MyDb.Query("select * from member")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(memberList)

}
