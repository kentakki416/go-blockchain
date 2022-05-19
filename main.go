package main

import (
	"fmt"
	"go-blockchain/wallet"
)

func main() {
	w := wallet.NewWallet()
	fmt.Println(w.PrivateKey())
	fmt.Println(w.PublicKey())
	fmt.Println(w.BlockchainAddress())
}
