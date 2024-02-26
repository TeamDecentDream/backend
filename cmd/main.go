package main

import (
	"backend/internal/config"
	"backend/internal/db"
	"log"
)

func main() {
	err := db.SetMySQL()
	if err != nil {
		log.Fatal("DB Connecting Error")
	}

	//err = config.SetEthSepoliaNet()
	//if err != nil {
	//	log.Fatal("Sepolia Error")
	//}

	router := config.SetRouter()
	err = router.Run(":8080")
	if err != nil {
		log.Fatal("Fail Opening the Server")
	}

}
