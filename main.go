package main

import (
	"go-blockchain/block"
	"go-blockchain/message"
)


func main() {
	messageM := message.NewMessage()
	messageA := message.NewMessage()
	messageB := message.NewMessage()
	// コンセンサスアルゴリズム未実装
	//tr := message.NewTransaction(messageA.PrivateKey(), messageA.PublicKey(), messageA.BlockchainAddress(), messageB.BlockchainAddress(), "Hello")

	blockchain := block.NewBlockchain(messageM.BlockchainAddress())
	blockchain.AddTransaction(messageA.BlockchainAddress(), messageB.BlockchainAddress(), "hello")
	blockchain.AddTransaction(messageB.BlockchainAddress(), messageA.BlockchainAddress(), "Hi !")

	blockchain.Mining()
	blockchain.Puts()

	blockchain.AddTransaction(messageA.BlockchainAddress(), messageB.BlockchainAddress(), "How are you?")
	blockchain.Mining()
	blockchain.Puts()

	blockchain.AddTransaction(messageB.BlockchainAddress(), messageA.BlockchainAddress(), "I'm fine!")
	blockchain.Puts()

	blockchain.Mining()
	blockchain.Puts()

	blockchain.PutsMessageList(messageA.BlockchainAddress())
	blockchain.PutsMessageList(messageB.BlockchainAddress())
}