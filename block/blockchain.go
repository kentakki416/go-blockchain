package main

import (
	"fmt"
	"time"
)

// ブロック構造体
type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

// ブロックの生成
func NewBlock(nonce int, previousHash string) *Block {
	b := new(Block)
	b.nonce = nonce
	b.previousHash = previousHash
	b.timestamp = time.Now().UnixNano()
	return b
}

// ブロックを出力
func (b *Block) Print() {
	fmt.Printf("nonce          %d\n", b.nonce)
	fmt.Printf("previous_hash  %s\n", b.previousHash)
	fmt.Printf("timestamp      %d\n", b.timestamp)
	fmt.Printf("transactions   %s\n", b.transactions)
}

func main() {
	b := NewBlock(0, "init hash")
	b.Print()
}
