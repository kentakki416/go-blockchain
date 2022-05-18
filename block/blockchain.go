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
	transactions []string
}

// ブロックの生成
func NewBlock(nonce int, previousHash [32]byte) *Block {
	b := new(Block)
	b.nonce = nonce
	b.previousHash = previousHash
	b.timestamp = time.Now().UnixNano()
	return b
}

// ブロックを出力
func (b *Block) Print() {
	fmt.Printf("nonce          %d\n", b.nonce)
	fmt.Printf("previous_hash  %x\n", b.previousHash)
	fmt.Printf("timestamp      %d\n", b.timestamp)
	fmt.Printf("transactions   %s\n", b.transactions)
}

// Hashの生成
func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256([]byte(m))
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Nonce        int      `json:nonce`
		PreviousHash [32]byte `json:previous_hash`
		Timestamp    int64    `json:timestamp`
		Transactions []string `json:transactions`
	}{
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Timestamp:    b.timestamp,
		Transactions: b.transactions,
	})
}

// ------------------------------------------------------------------------------------------
type Blockchain struct {
	transactionPool []string
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
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
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

func main() {
	blockChain := NewBlockchain()

	// 最後のブロックのハッシュを生成
	previousHash := blockChain.LastBlock().Hash()
	blockChain.CreateBlock(5, previousHash)

	previousHash = blockChain.LastBlock().Hash()
	blockChain.CreateBlock(2, previousHash)
	blockChain.Print()
}
