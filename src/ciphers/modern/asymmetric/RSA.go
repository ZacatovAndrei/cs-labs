package asymmetric

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/hex"
	"errors"
	"math/big"
)

type RSA struct {
	privateKey *big.Int
	publicKey  *big.Int
	modulus    *big.Int
}

func NewRSA() *RSA {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	return &RSA{
		privateKey: key.D,
		publicKey:  big.NewInt(int64(key.E)),
		modulus:    key.N,
	}
}

func (R RSA) GetType() string {
	return "RSA asymmetrical cipher"
}

func (R RSA) Encode(plain string) string {
	if len(plain) > (2048 / 8) {
		panic(errors.New("text too long to encode"))
	}
	msg := big.NewInt(0)
	msg.SetBytes([]byte(plain))
	return hex.EncodeToString(msg.Exp(msg, R.publicKey, R.modulus).Bytes())
}

func (R RSA) Decode(cipher string) string {
	dCipher, err := hex.DecodeString(cipher)
	if err != nil {
		panic(err)
	}
	msg := big.NewInt(0)
	msg.SetBytes(dCipher)
	return string(msg.Exp(msg, R.privateKey, R.modulus).Bytes())
}
