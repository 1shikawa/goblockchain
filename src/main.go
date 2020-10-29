package main

import (
	"fmt"
	"lock"
	"log"
	"wallet"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	// w := wallet.NewWallet()
	// fmt.Println(w.PrivateKeyStr())
	// fmt.Println(w.PublicKeyStr())
	// fmt.Println(w.BlockchainAddress())

	// t := wallet.NewTransaction(w.PrivateKey(), w.PublicKey(), w.BlockchainAddress(), "B", 1.0)
	// fmt.Printf("signature %s\n", t.GenerateSignature())

	walletM := wallet.NewWallet()
	walletA := wallet.NewWallet()
	walletB := wallet.NewWallet()

	// Wallet
	t := wallet.NewTransaction(walletA.PrivateKey(), walletA.PublicKey(), walletA.BlockchainAddress(), walletB.BlockchainAddress(), 1.0)

	// Blockchain
	blockchain := block.NewBlockchain(walletM.BlockchainAddress())
	isAdded := blockchain.AddTransaction(walletA.BlockchainAddress(), walletB.BlockchainAddress(), 1.0,
		walletA.PublicKey(), t.GenerateSignature())
	fmt.Println("Added? ", isAdded)

	blockchain.Mining()
	blockchain.Print()

	fmt.Printf("A %.1f\n", blockchain.CalculateTotalAmount(walletA.BlockchainAddress()))
	fmt.Printf("B %.1f\n", blockchain.CalculateTotalAmount(walletB.BlockchainAddress()))
	fmt.Printf("M %.1f\n", blockchain.CalculateTotalAmount(walletM.BlockchainAddress()))

}
