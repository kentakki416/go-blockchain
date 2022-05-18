package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// ブロック構造体
type Block struct {
	nonce        int
	previousHash [32]byte
	timestamp    int64
	transactions []*Transaction
}

// ブロックの生成
func NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {
	b := new(Block)
	b.nonce = nonce
	b.previousHash = previousHash
	b.timestamp = time.Now().UnixNano()
	b.transactions = transactions
	return b
}

// ブロックを出力
func (b *Block) Print() {
	fmt.Printf("nonce          %d\n", b.nonce)
	fmt.Printf("previous_hash  %x\n", b.previousHash)
	fmt.Printf("timestamp      %d\n", b.timestamp)
	for _, t := range b.transactions {
		t.Print()
	}
}

// Hashの生成
func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256([]byte(m))
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Nonce        int            `json:nonce`
		PreviousHash [32]byte       `json:previous_hash`
		Timestamp    int64          `json:timestamp`
		Transactions []*Transaction `json:transactions`
	}{
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Timestamp:    b.timestamp,
		Transactions: b.transactions,
	})
}

// ------------------------------------------------------------------------------------------
type Blockchain struct {
	transactionPool []*Transaction
	chain           []*Block
}

// ブロックチェーンの作成
func NewBlockchain() *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.CreateBlock(0, b.Hash())
	return bc
}

// ブロックチェーンの中にブロックを格納
func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	// ブロックをブロックチェーンに追加した際にPoolを空にする
	bc.transactionPool = []*Transaction{}
	return b
}

// ブロックチェーンの中の最後のブロックを取得
func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

// ブロッックチェーンの出力
func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

// TransactionPoolにTransactionを追加
func (bc *Blockchain) AddTransaction(sender string, recipient string, value float32) {
	t := NewTransaction(sender, recipient, value)
	bc.transactionPool = append(bc.transactionPool, t)
}

// ------------------------------------------------------------------------------------------------
type Transaction struct {
	senderBLockchainAddress    string
	recipientBlockchainAddress string
	value                      float32
}

// Transactionの生成
func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}

// Transactionの出力
func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf(" sender_blockchain_address    %s\n", t.senderBLockchainAddress)
	fmt.Printf(" recipient_blockchain_address %s\n", t.recipientBlockchainAddress)
	fmt.Printf(" value                        %.1f\n", t.value)
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender_blockchain_address"`
		Recipient string  `json:"recipient_blockchain_address"`
		Value     float32 `json:"value"`
	}{
		Sender:    t.senderBLockchainAddress,
		Recipient: t.recipientBlockchainAddress,
		Value:     t.value,
	})
}

func main() {
	blockChain := NewBlockchain()

	blockChain.AddTransaction("A", "B", 10)
	// 最後のブロックのハッシュを生成
	previousHash := blockChain.LastBlock().Hash()
	blockChain.CreateBlock(5, previousHash)

	previousHash = blockChain.LastBlock().Hash()
	blockChain.CreateBlock(2, previousHash)
	blockChain.Print()
}
