package block

import (
	"fmt"
	"strings"
	"encoding/json"
)

type Transaction struct {
	sendAddress string
	receiveAddress string
	message string
}

func NewTransaction(sender string, recipient string, message string) *Transaction {
	return &Transaction{sender, recipient, message}
}

func (t *Transaction) Puts(i int) {
	fmt.Printf("%s Transaction No.%d %s\n", strings.Repeat("-", 23),i,strings.Repeat("-", 23))
	fmt.Printf(" sender_blockchain_address      %s\n", t.sendAddress)
	fmt.Printf(" recipient_blockchain_address   %s\n", t.receiveAddress)
	fmt.Printf(" message                        %s\n", t.message)
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender_blockchain_address"`
		Recipient string  `json:"recipient_blockchain_address"`
		Message     string `json:"value"`
	}{
		Sender:    t.sendAddress,
		Recipient: t.receiveAddress,
		Message:     t.message,
	})
}

