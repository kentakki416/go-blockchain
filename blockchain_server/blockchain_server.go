package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

type BlockchainServer struct {
	port uint16
}

// ブロックチェーンサーバーの作成
func NewBlockchainServer(port uint16) *BlockchainServer {
	return &BlockchainServer{port}
}

// ブロックチェーンサーバーのポートを返す
func (bcs *BlockchainServer) Port() uint16 {
	return bcs.port
}

// httpサーバーのハンドル
func HelloWorld(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello World!")
}

// サーバーの立ち上げ
func (bcs *BlockchainServer) Run() {
	http.HandleFunc("/", HelloWorld)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(bcs.Port())), nil))
}
