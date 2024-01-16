package main

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
)

var infuraURL = ""

func main() {
	ethclient.DialContext(context.Background(), infuraURL)
}
