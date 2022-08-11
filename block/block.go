package block

import (
	"fmt"
	"time"
	"encoding/json"
	"crypto/sha256"
)

type Block struct {
	timestamp      int64
	preHash        [32]byte
	transactions   []*Transaction
	nonce          int
}


func NewBlock(nonce int, preHash [32]byte, transactions []*Transaction, )(*Block){
	var block Block
	block.timestamp = time.Now().UnixNano()
	block.preHash = preHash
	block.transactions = transactions
	block.nonce = nonce
	return &block
}

func (bl *Block) Puts(){
	fmt.Printf("timestamp       %d\n", bl.timestamp)
	fmt.Printf("nonce           %d\n", bl.nonce)
	fmt.Printf("previous_hash   %x\n", bl.preHash)
	for i, t := range bl.transactions {
		t.Puts(i)
	}
}

func (bl *Block) Hash() [32]byte {
	m, _ := json.Marshal(bl)
	return sha256.Sum256([]byte(m))
}

func (bl *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64          `json:"timestamp"`
		Nonce        int            `json:"nonce"`
		PreHash      string         `json:"pre_hash"`
		Transactions []*Transaction `json:"transactions"`
	}{
		Timestamp:    bl.timestamp,
		Nonce:        bl.nonce,
		PreHash:      fmt.Sprintf("%x", bl.preHash),
		Transactions: bl.transactions,
	})
}
