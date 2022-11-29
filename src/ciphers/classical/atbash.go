package classical

import (
	"strings"
	"unicode"
)

type Atbash struct {
	alphabet string
}

func NewAtbash(alphabet string) *Atbash {
	return &Atbash{alphabet: alphabet}
}

func (A Atbash) Encode(plainText string) string {
	// setup
	plainText = strings.ToUpper(plainText)
	var res strings.Builder
	var newIndex int
	//Encryption
	for _, char := range plainText {
		if !unicode.IsLetter(char) {
			res.WriteRune(char)
			continue
		}
		newIndex = len(A.alphabet) - strings.Index(A.alphabet, string(char)) - 1
		res.WriteRune(rune(A.alphabet[newIndex]))
	}

	return res.String()
}

func (A Atbash) Decode(cipherText string) string {
	res := A.Encode(cipherText)
	res = strings.ToLower(res)
	return res
}

func (A Atbash) GetType() string {
	return "Atbash Cipher"
}
