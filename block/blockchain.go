package block

import (
	"fmt"
	"strings"
	"encoding/json"
)

const (
	DIFF = 4
)

type Blockchain struct {
	current_transactions []*Transaction
	chain []*Block
	blockchainAddress string
}

func NewBlockchain(blockchainAddress string) *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.blockchainAddress = blockchainAddress
	bc.CreateBlock(0, b.Hash())
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash, bc.current_transactions)
	bc.chain = append(bc.chain, b)
	bc.current_transactions = []*Transaction{}
	return b
}

func (bc *Blockchain) CurrentTransactions() []*Transaction {
	return bc.current_transactions
}

func (bc *Blockchain) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Blocks []*Block `json:"blockchains"`
	}{
		Blocks: bc.chain,
	})
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc *Blockchain) Puts() {
	for i, block := range bc.chain {
		fmt.Printf("%s Block No.%d %s\n", strings.Repeat("~", 25), i, strings.Repeat("~", 25))
		block.Puts()
	}
	fmt.Printf("%s end of Block %s\n", strings.Repeat("~", 25), strings.Repeat("~", 25))
}

func (bc *Blockchain) AddTransaction(sender string, recipient string, message string) bool {
	t := NewTransaction(sender, recipient, message)
	bc.current_transactions = append(bc.current_transactions, t)
	return true
}

func (bc *Blockchain) Mining() bool {
	nonce := bc.ProofOfWork()
	preHash := bc.LastBlock().Hash()
	bc.CreateBlock(nonce, preHash)
	return true
}

func (bc *Blockchain) ProofOfWork() int {
	transactions := bc.CopyCurrentTransactions()
	previousHash := bc.LastBlock().Hash()
	nonce := 0
	for !bc.ValidProof(nonce, previousHash, transactions, DIFF) {
		nonce += 1
	}
	return nonce
}

func (bc *Blockchain) CopyCurrentTransactions() []*Transaction {
	transactions := make([]*Transaction, 0)
	for _, t := range bc.current_transactions {
		transactions = append(transactions,NewTransaction(t.sendAddress,
														  t.receiveAddress,
														  t.message))
	}
	return transactions
}

func (bc *Blockchain) ValidProof(nonce int, preHash [32]byte, transactions []*Transaction, difficulty int) bool {
	zeros := strings.Repeat("0", difficulty)
	guessBlock := Block{0, preHash, transactions, nonce}
	guessHashStr := fmt.Sprintf("%x", guessBlock.Hash())
	return guessHashStr[:difficulty] == zeros
}

func (bc *Blockchain) MessageList(blockchainAddress string) []string {
	messageList := make([]string,0)
	for _, b := range bc.chain {
		for _, t := range b.transactions {
			message := t.sendAddress + " : " +t.message
			if blockchainAddress == t.receiveAddress {
				messageList = append(messageList, message)
			}
		}
	}
	return messageList
}

func (bc *Blockchain) PutsMessageList(blockchainAdress string) {
	messageList := bc.MessageList(blockchainAdress)
	fmt.Printf("%s %s's message list %s\n", strings.Repeat("=", 5), blockchainAdress ,strings.Repeat("=", 5))
	for _ , s := range messageList{
		fmt.Println(s)
	} 
}