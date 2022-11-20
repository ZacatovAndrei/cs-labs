package classical

import (
	"strings"
	"unicode"
)

type Vigenere struct {
	alphabet string
	key      string
}

func (v Vigenere) GetType() string {
	return "Vigenere Cipher"
}

func (v Vigenere) Encode(plain string) string {
	plain = strings.ToUpper(plain)
	var (
		keyPos   = 0
		newIndex = 0
		res      strings.Builder
	)
	//Encryption
	for _, char := range plain {
		if !unicode.IsLetter(char) {
			res.WriteRune(char)
			continue
		}
		newIndex = strings.Index(v.alphabet, string(char)) + strings.Index(v.alphabet, string(v.key[keyPos])) + len(v.alphabet)
		newIndex %= len(v.alphabet)
		keyPos = (keyPos + 1) % len(v.key)
		res.WriteRune(rune(v.alphabet[newIndex]))
	}
	return res.String()

}

func (v Vigenere) Decode(cipher string) string {
	return "decoded"
	panic("implement me")
}

func NewVigenere(alphabet string, key string) *Vigenere {
	return &Vigenere{alphabet: alphabet, key: strings.ToUpper(key)}
}
