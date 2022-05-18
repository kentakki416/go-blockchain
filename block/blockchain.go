package main

import (
	"fmt"
	"strings"
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

// ------------------------------------------------------------------------------------------
type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

// ブロックチェーンの作成
func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "Init hash")
	return bc
}

// ブロックチェーンの中にブロックを格納
func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
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
	blockChain.Print()
	blockChain.CreateBlock(5, "hash 1")
	blockChain.Print()
	blockChain.CreateBlock(2, "hash 2")
	blockChain.Print()
}
