package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

type Wallet struct {
	privateKey        *ecdsa.PrivateKey
	publicKey         *ecdsa.PublicKey
	blockchainAddress string
}

// Walletの新規作成
func NewWallet() *Wallet {
	// 1. Createing ECDSA privateKey & publicKey
	w := new(Wallet)
	// pricvateKeyの作成
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	w.privateKey = privateKey
	w.publicKey = &w.privateKey.PublicKey

	// 2. Perform SHA-256 hashing on the publicKey
	h2 := sha256.New()
	h2.Write(w.publicKey.X.Bytes())
	h2.Write(w.publicKey.Y.Bytes())
	digest2 := h2.Sum(nil)
	// 3. Perform RIPEMD-160 hashing on the result of SHA-256
	h3 := ripemd160.New()
	h3.Write(digest2)
	digest3 := h3.Sum(nil)
	// 4. Add version byte in front fo RIPEMD-160hash (0x00 for Main Nerwork)
	vd4 := make([]byte, 21)
	vd4[0] = 0x00
	copy(vd4[1:], digest3[:])
	// 5. Perform SHA-256 hash on the extended RIPEMD-160 result
	h5 := sha256.New()
	h5.Write(vd4)
	digest5 := h5.Sum(nil)
	// 6. Perform SHA-256 hash on the result of the previous SHA-256 hash
	h6 := sha256.New()
	h6.Write(digest5)
	digest6 := h6.Sum(nil)
	// 7. Take the first 4 bytes form 7 at the end fo extended RIPMED-160 hash from 4
	chsum := digest6[:4]
	// 8. Add the 4 checksum bytes from 7 at the end fo extended RIPMED-160 hash from 4
	dc8 := make([]byte, 25)
	copy(dc8[:21], vd4[:])
	copy(dc8[21:], chsum[:])
	// 9. Convert the result from a byte string into base58
	address := base58.Encode(dc8)
	w.blockchainAddress = address

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

func (w *Wallet) BlockchainAddress() string {
	return w.blockchainAddress
}
