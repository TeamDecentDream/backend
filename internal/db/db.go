package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var MyDb *sql.DB

func SetMySQL() error {
	db, err := sql.Open("mysql", "nextfarm:nextfarm@tcp(127.0.0.1:3306)/nextfarm?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
		return err
	} else {
		MyDb = db
		log.Println(MyDb)
		return nil
	}
}
