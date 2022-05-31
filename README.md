# Go-blockchain
 
ブロックチェーンを用いた仮想通貨の送金システムデモアプリ

ウォレットサーバーとブロックチェーンネットワークを構築し、仮想通貨を送受信できる
# DEMO
 
![アプリの概要図](images/demo.png)
![アプリの画面](images/screen.png)
  
# Requirement
 
"Go-blockchain"を動かすのに必要なライブラリなどを列挙する
 
* go 1.18
* macOS(windowでの検証はしてない)
* tailwindcss
 
# Installation
 
必要なライブラリのインストール
 
```bash
go get github.com/btcsuite/btcutil/base58
go get golang.org/x/crypto/ripemd160
```
 
# Usage
 
アプリの実行方法
 
```bash
git clone https://github.com/kentakki416/go-blockchain.git

cd go-blockchain

go mod init go-blockchain　(vscodeの場合)

（上記のライブラリをインストール）

cd wallet_server

// Aさんのウォレットサーバー立ち上げ
go run main.go wallet_server.go

// Bさんのウォレットサーバー立ち上げ(新しいターミナルで)
go run main.go wallet_server.go -port 8081 -gateway http://127.0.0.1:5001

cd blockchain_server

// ブロックチェーンサーバー(port:5001)の立ち上げ（新しいターミナルで）
go run main.go blockchain_server.go

// ブロックチェーンサーバー(port:5002)の立ち上げ（新しいターミナルで）
go run main.go blockchain_server.go -port 5002

// ブロックチェーンサーバー(port:5003)の立ち上げ（新しいターミナルで）
go run main.go blockchain_server.go -port 5003

```

# Function
実装した機能の概要紹介
* ブロックチェーンの生成
* トランザクションの生成
* マイニング
* 他のブロックチェーンサーバと通信
* コンセンサスアルゴリズム
* ブロックの伝播

# Note
 
注意点などがあれば書く
 
 
# License 
"Go-blockchain" is under [MIT license](https://en.wikipedia.org/wiki/MIT_License).
 
