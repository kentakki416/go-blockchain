package main

import "time"

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
	b.timestamp = time.Now().UnixNano()
	b.previousHash = previousHash
	return b
}

func main() {
	b := NewBlock(0, "init hash")
	fmt.Pringln(b)
}
