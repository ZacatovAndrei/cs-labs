package signature

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/hex"
	"fmt"
	"hash"
	"math/big"
)

type Signature struct {
	digest hash.Hash
	prKey  *big.Int
	PbKey  *big.Int
	Mod    *big.Int
}

func NewSignature(digest hash.Hash) *Signature {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	return &Signature{
		digest: digest,
		prKey:  key.D,
		PbKey:  big.NewInt(int64(key.E)),
		Mod:    key.N,
	}
}
func (s Signature) Sign(message string) (string, string) {
	msg := []byte(message)
	s.digest.Write(msg)
	// Getting the hash sum
	hashSum := s.digest.Sum(nil)
	msgNum := big.NewInt(0)
	msgNum.SetBytes(hashSum)
	out := msgNum.Exp(msgNum, s.prKey, s.Mod)
	return out.Text(16), hex.EncodeToString(hashSum)
}

func (s Signature) Verify(message string, original string, publicKey *big.Int, modulus *big.Int) bool {
	msgNum := big.NewInt(0)
	msgNum.SetString(message, 16)

	origMsg := big.NewInt(0)
	origMsg.SetString(original, 16)

	res := msgNum.Exp(msgNum, publicKey, modulus)

	fmt.Println(res)
	fmt.Println(origMsg)
	if res.Cmp(origMsg) != 0 {
		return false
	}
	return true
}
