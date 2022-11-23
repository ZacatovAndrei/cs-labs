package asymmetric

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/hex"
	"errors"
	"log"
	"math/big"
)

type RSA struct {
	privateKey *big.Int
	PublicKey  *big.Int
	Modulus    *big.Int
}

func NewRSA() *RSA {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	log.Printf("Generating RSA Keypair\n\nModulus: %v\nPublic key: %v\n\n", key.N.String(), key.E)
	return &RSA{
		privateKey: key.D,
		PublicKey:  big.NewInt(int64(key.E)),
		Modulus:    key.N,
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
	return msg.Exp(msg, R.PublicKey, R.Modulus).Text(16)
}

func (R RSA) Decode(cipher string) string {
	dCipher, err := hex.DecodeString(cipher)
	if err != nil {
		panic(err)
	}
	if len(dCipher) > (2048 / 8) {
		panic("value too big to decode")
	}
	msg := big.NewInt(0)
	msg.SetBytes(dCipher)
	return string(msg.Exp(msg, R.privateKey, R.Modulus).Bytes())
}
