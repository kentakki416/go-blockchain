package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
)

type Wallet struct {
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
}

// Walletの新規作成
func NewWallet() *Wallet {
	w := new(Wallet)
	// pricvateKeyの作成
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	w.privateKey = privateKey
	w.publicKey = &w.privateKey.PublicKey
	return w
}

// privateKeyの取得
func (w *Wallet) PrivateKey() *ecdsa.PrivateKey {
	return w.privateKey
}

// PrivateKeyの中身を出力
func (w *Wallet) PrivateKeyStr() string {
	return fmt.Sprintf("%x", w.privateKey.D.Bytes())
}

// publicKeyの取得
func (w *Wallet) PublicKey() *ecdsa.PublicKey {
	return w.publicKey
}

// PublicKeyの中身を出力
func (w *Wallet) PublicKeyStr() string {
	return fmt.Sprintf("%x%x", w.publicKey.X.Bytes(), w.publicKey.Y.Bytes())
}
