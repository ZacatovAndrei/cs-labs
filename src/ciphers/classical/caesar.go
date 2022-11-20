package classical

import (
	"strconv"
	"strings"
	"unicode"
)

type Caesar struct {
	alphabet string
	key      string
}

func NewCaesar(alphabet string, key string) *Caesar {
	return &Caesar{alphabet: alphabet, key: key}
}

func (c Caesar) GetType() string {
	return "Caesar's cipher"
}

func (c Caesar) Encode(plain string) string {
	// setup
	shift, _ := strconv.ParseInt(c.key, 10, 64)
	shift %= int64(len(c.alphabet))
	plain = strings.ToUpper(plain)
	var res strings.Builder
	var newIndex int
	//Encryption
	for _, char := range plain {
		if !unicode.IsLetter(char) {
			res.WriteRune(char)
			continue
		}
		newIndex = strings.Index(c.alphabet, string(char)) + int(shift) + len(c.alphabet)
		newIndex %= len(c.alphabet)
		res.WriteRune(rune(c.alphabet[newIndex]))
	}
	return res.String()
}

func (c Caesar) Decode(cipher string) string {
	// setup
	shift, _ := strconv.ParseInt(c.key, 10, 64)
	shift %= int64(len(c.alphabet))
	cipher = strings.ToUpper(cipher)
	var res strings.Builder
	var newIndex int
	//Encryption
	for _, char := range cipher {
		if !unicode.IsLetter(char) {
			res.WriteRune(char)
			continue
		}
		newIndex = strings.Index(c.alphabet, string(char)) - int(shift) + len(c.alphabet)
		newIndex %= len(c.alphabet)
		res.WriteRune(rune(c.alphabet[newIndex]))
	}
	return strings.ToLower(res.String())
}
