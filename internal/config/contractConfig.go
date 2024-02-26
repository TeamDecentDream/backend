package config

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
)

var Client *ethclient.Client

func SetEthSepoliaNet() error {
	var err error
	Client, err = ethclient.Dial("인퓨라 세폴리아 키!")
	if err != nil {
		fmt.Println("client error")
		return err
	}
	return nil
}
