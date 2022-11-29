package classical

import (
	"strings"
)

type CaesarWithPermutation struct {
	permutation string
	cipher      Caesar
}

func (cp CaesarWithPermutation) GetType() string {
	return cp.cipher.GetType() + " with permutation"
}

func (cp CaesarWithPermutation) Encode(plainText string) string {
	return cp.cipher.Encode(plainText)
}

func (cp CaesarWithPermutation) Decode(cipherText string) string {
	return cp.cipher.Decode(cipherText)
}

func NewCaesarWithPermutation(alphabet string, key string, permutation string) *CaesarWithPermutation {
	alphabet = permuteAlphabet(alphabet, permutation)
	inner := *NewCaesar(alphabet, key)
	return &CaesarWithPermutation{cipher: inner, permutation: permutation}
}

func permuteAlphabet(alphabet string, permutation string) string {
	//TODO make a check to see if the alphabets are incompatible
	permutation = strings.ToUpper(permutation)
	var res strings.Builder
	var letters = make(map[rune]bool, len(alphabet))
	for _, char := range alphabet {
		letters[char] = true
	}
	// filing the string builder with initial alphabet
	for _, char := range permutation {
		if r, ok := letters[char]; ok && r {
			res.WriteRune(char)
			letters[char] = false
		}
	}

	for _, char := range alphabet {
		if r, ok := letters[char]; r && ok {
			res.WriteRune(char)
		}
	}
	return res.String()
}
