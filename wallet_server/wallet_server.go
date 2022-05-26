package main

import (
	"bytes"
	"encoding/json"
	"go-blockchain/block"
	"go-blockchain/utils"
	"go-blockchain/wallet"
	"html/template"
	"io"
	"log"
	"net/http"
	"path"
	"strconv"
)

const tempDir = "templates"

// wallet（ユーザー）の構造た
type WalletServer struct {
	port    uint16
	gateway string // 接続するBlockchainNode
}

// Walletの作成
func NewWalletServer(port uint16, gateway string) *WalletServer {
	return &WalletServer{port: port, gateway: gateway}
}

func (ws *WalletServer) Port() uint16 {
	return ws.port
}

func (ws *WalletServer) Gateway() string {
	return ws.gateway
}

// indexページを表示するハンドル
func (ws *WalletServer) Index(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		// htmlファイルを表示
		t, _ := template.ParseFiles(path.Join(tempDir, "index.html"))
		t.Execute(w, "")
	default:
		log.Printf("ERROR: Invalid HTTP Method")
	}
}

// ウォレットを作成するAPIハンドル
func (ws *WalletServer) Wallet(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		w.Header().Add("Content-Type", "application/json")
		myWallet := wallet.NewWallet()
		m, _ := myWallet.MarshalJSON()
		io.WriteString(w, string(m[:]))
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: Invalid HTTP Method")
	}
}

// トランザクションを作成して、ブロックチェーンネットワークに送信するAPIハンドル
func (ws *WalletServer) CreateTransaction(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		// フロントからくるJsonをStructに格納する
		decoder := json.NewDecoder(req.Body)
		var t wallet.TransactionRequest
		err := decoder.Decode(&t)
		if err != nil {
			log.Printf("ERROR: %v ", err)
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}
		// Jsonのバリデーション処理
		if !t.Validate() {
			log.Println("ERROR: missing field(s)")
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}

		// publicKeyとprivateKeyを生成
		publicKey := utils.PublicKeyFromString(*t.SenderPublicKey)
		privateKey := utils.PrivateKeyFromString(*t.SenderPrivateKey, publicKey)
		// valueを生成
		value, err := strconv.ParseFloat(*t.Value, 32)
		if err != nil {
			log.Println("ERROR: parse error")
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}
		value32 := float32(value)

		w.Header().Add("Content-Type", "application/json")
		// transactionの生成
		transaction := wallet.NewTransaction(privateKey, publicKey, *t.SenderBlockchainAddress,
			*t.RecipientBlockchainAddress, value32)
		// signatureの生成
		signature := transaction.GenerateSignature()
		signatureStr := signature.String()

		bt := &block.TransactionRequest{
			t.SenderBlockchainAddress,
			t.RecipientBlockchainAddress,
			t.SenderPublicKey,
			&value32, &signatureStr,
		}
		m, _ := json.Marshal(bt)
		buf := bytes.NewBuffer(m)

		resp, _ := http.Post(ws.Gateway()+"/transactions", "application/json", buf)

		if resp.StatusCode == 201 {
			io.WriteString(w, string(utils.JsonStatus("success")))
			return
		}
		io.WriteString(w, string(utils.JsonStatus("fail")))

	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: Invalid HTTP Method")
	}
}

func (ws *WalletServer) Run() {
	http.HandleFunc("/", ws.Index)
	http.HandleFunc("/wallet", ws.Wallet)
	http.HandleFunc("/transaction", ws.CreateTransaction)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(ws.Port())), nil))
}
