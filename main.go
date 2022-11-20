package main

import (
	"cs-labs/src/ciphers"
	"cs-labs/src/ciphers/classical"
	"fmt"
)

const (
	EngAlph = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Msg     = "This is some Plaintext. Also, the classical ciphers loose all case-sensitive Data with classical ciphers."
)

func init() {
	fmt.Printf("This is a sample driver program to Show the results of the CS laboratory works\n")
	fmt.Printf("Author:\nGroup:\n")
	fmt.Printf("Test message is: %v\n\n", Msg)
}

func main() {
	var ciphersImplemented []ciphers.Cipher
	// Adding the implemented ciphers and checking them via encryption and decryption
	ciphersImplemented = append(ciphersImplemented,
		classical.NewAtbash(EngAlph))
	ciphersImplemented = append(ciphersImplemented,
		classical.NewCaesar(EngAlph, "14"))
	ciphersImplemented = append(ciphersImplemented,
		classical.NewCaesarWithPermutation(EngAlph, "13", "This is a test"))
	ciphersImplemented = append(ciphersImplemented,
		classical.NewVigenere(EngAlph, "Attack"))
	// ciphersImplemented = append(ciphersImplemented,)
	// ciphersImplemented = append(ciphersImplemented,)
	// ciphersImplemented = append(ciphersImplemented,)
	// ciphersImplemented = append(ciphersImplemented,)
	for _, cipher := range ciphersImplemented {
		encoded := cipher.Encode(Msg)
		fmt.Printf("Current cipher: %v\nEncoded message: %v\nDecoded message: %v\n", cipher.GetType(), encoded, cipher.Decode(encoded))
	}

}
