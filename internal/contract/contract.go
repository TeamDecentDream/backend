package contract

import (
	"backend/internal/config"
	NZFToken "backend/internal/contract"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"time"
)

func loadContract(address_ string) {
	privateKey, err := crypto.HexToECDSA("프라이빗키~")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(address_)
	instance, err := NZFToken.NewNZFToken(address, config.Client)
	auth := bind.NewKeyedTransactor(privateKey)
	poolAddress := common.HexToAddress("Defi Pool Address!")
	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			_, err := instance.BuyBack(auth, poolAddress, 0)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
