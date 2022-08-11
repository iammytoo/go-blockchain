package message

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type Message struct {
	privateKey        *ecdsa.PrivateKey
	publicKey         *ecdsa.PublicKey
	blockchainAddress string
}

func NewMessage() *Message {
	m := new(Message)
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	m.privateKey = privateKey
	m.publicKey = &m.privateKey.PublicKey
	privateKey, err := crypto.GenerateKey()
    if err != nil {
        log.Fatal(err)
    }
    privateKeyBytes := crypto.FromECDSA(privateKey)
    fmt.Println("SAVE BUT DO NOT SHARE THIS (Private Key):", hexutil.Encode(privateKeyBytes))
    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    }
    publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
    fmt.Println("Public Key:", hexutil.Encode(publicKeyBytes)) 
    address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
    fmt.Println("Address:", address)
	m.blockchainAddress = address
	return m
}

func (m *Message) PrivateKey() *ecdsa.PrivateKey {
	return m.privateKey
}

func (m *Message) PrivateKeyStr() string {
	return fmt.Sprintf("%x", m.privateKey.D.Bytes())
}

func (m *Message) PublicKey() *ecdsa.PublicKey {
	return m.publicKey
}

func (m *Message) PublicKeyStr() string {
	return fmt.Sprintf("%064x%064x", m.publicKey.X.Bytes(), m.publicKey.Y.Bytes())
}

func (m *Message) BlockchainAddress() string {
	return m.blockchainAddress
}


func (m *Message) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		PrivateKey        string `json:"private_key"`
		PublicKey         string `json:"public_key"`
		BlockchainAddress string `json:"blockchain_address"`
	}{
		PrivateKey:        m.PrivateKeyStr(),
		PublicKey:         m.PublicKeyStr(),
		BlockchainAddress: m.BlockchainAddress(),
	})
}

type Transaction struct {
	senderPrivateKey      *ecdsa.PrivateKey
	senderPublicKey       *ecdsa.PublicKey
	sendAddress           string
	receiveAddress        string
	message               string
}

func NewTransaction(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey,
	sender string, recipient string,message string) *Transaction {
	return &Transaction{privateKey, publicKey, sender, recipient, message}
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender_blockchain_address"`
		Recipient string  `json:"recipient_blockchain_address"`
		Message   string `json:"message"`
	}{
		Sender:    t.sendAddress,
		Recipient: t.receiveAddress,
		Message:   t.message,
	})
}
