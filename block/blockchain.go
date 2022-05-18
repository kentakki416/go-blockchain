package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

const (
	MINING_DIFFICULTY = 3
	MINING_SENDER     = "THE BLOCKCHAIN"
	MINING_REWARD     = 1.0
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
	transactionPool   []*Transaction
	chain             []*Block
	blockchainAddress string //ブロックチェーンネットワークを構成する各nodeのアドレス
}

// ブロックチェーンの作成
func NewBlockchain(blockchainAddress string) *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.blockchainAddress = blockchainAddress
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

// Poolに溜まっているTransactionをコピーする
func (bc *Blockchain) CopyTransactionPool() []*Transaction {
	transactions := make([]*Transaction, 0)
	for _, t := range bc.transactionPool {
		transactions = append(transactions,
			NewTransaction(t.senderBlockchainAddress, t.recipientBlockchainAddress, t.value),
		)
	}
	return transactions
}

// NonceとpreviousHashとtransactionを使ってDifficultyを求める
func (bc *Blockchain) ValidProof(nonce int, previousHash [32]byte, transactions []*Transaction, difficulty int) bool {
	zeros := strings.Repeat("0", difficulty)
	guessBlock := Block{nonce, previousHash, 0, transactions}

	guessHashStr := fmt.Sprintf("%x", guessBlock.Hash())
	// 解の左から３つ目が000ならばtrue
	return guessHashStr[:difficulty] == zeros
}

// 解となるnonceを求める
func (bc *Blockchain) ProofOfWork() int {
	transactions := bc.CopyTransactionPool()
	previousHash := bc.LastBlock().Hash()
	nonce := 0
	for !bc.ValidProof(nonce, previousHash, transactions, MINING_DIFFICULTY) {
		nonce += 1
	}

	return nonce
}

// マイニング処理
func (bc *Blockchain) Mining() bool {
	// MINING_SENDERがbc.blockchainAddressにMINING_REWARD送るトランザクション
	bc.AddTransaction(MINING_SENDER, bc.blockchainAddress, MINING_REWARD)
	nonce := bc.ProofOfWork()
	previousHash := bc.LastBlock().Hash()
	bc.CreateBlock(nonce, previousHash)
	log.Println("action=mining, status=success")
	return true
}

// ユーザーのvalueの合計値を取得
func (bc *Blockchain) CalculateTotalAmount(blockchainAddress string) float32 {
	var totalAmount float32 = 0.0
	for _, b := range bc.chain {
		for _, t := range b.transactions {
			value := t.value
			// 受け取りの場合
			if blockchainAddress == t.recipientBlockchainAddress {
				totalAmount += value
			}
			if blockchainAddress == t.senderBlockchainAddress {
				totalAmount -= value
			}
		}
	}
	return totalAmount
}

// ------------------------------------------------------------------------------------------------
type Transaction struct {
	senderBlockchainAddress    string
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
	fmt.Printf(" sender_blockchain_address    %s\n", t.senderBlockchainAddress)
	fmt.Printf(" recipient_blockchain_address %s\n", t.recipientBlockchainAddress)
	fmt.Printf(" value                        %.1f\n", t.value)
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender_blockchain_address"`
		Recipient string  `json:"recipient_blockchain_address"`
		Value     float32 `json:"value"`
	}{
		Sender:    t.senderBlockchainAddress,
		Recipient: t.recipientBlockchainAddress,
		Value:     t.value,
	})
}

func main() {
	myBlockchainAddress := "my_blockchain_address"
	blockChain := NewBlockchain(myBlockchainAddress)

	blockChain.AddTransaction("A", "B", 10)
	blockChain.Mining()

	blockChain.AddTransaction("C", "D", 2)
	blockChain.AddTransaction("X", "Y", 5)
	blockChain.Mining()
	blockChain.Print()

	fmt.Printf("my %.1f\n", blockChain.CalculateTotalAmount("my_blockchain_address"))
	fmt.Printf("C %.1f\n", blockChain.CalculateTotalAmount("C"))
	fmt.Printf("D %.1f\n", blockChain.CalculateTotalAmount("D"))
}
