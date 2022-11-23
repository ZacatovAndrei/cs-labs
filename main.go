package main

import (
	"crypto/sha1"
	"cs-labs/src/ciphers"
	"cs-labs/src/ciphers/classical"
	"cs-labs/src/ciphers/modern/asymmetric"
	"cs-labs/src/ciphers/modern/symmetric"
	"cs-labs/src/signature"
	"fmt"
)

const (
	EngAlph = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Msg     = "This is some Plaintext. Also, the classical ciphers loose all case-sensitive Data with classical ciphers."
)

func init() {
	fmt.Printf("This is a sample driver program to Show the results of the CS laboratory works\n")
	fmt.Printf("Author:Zacatov Andrei\nGroup:FAF-201\n")
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
	ciphersImplemented = append(ciphersImplemented,
		symmetric.NewRC4("Some random key of a decent length"))
	ciphersImplemented = append(ciphersImplemented,
		symmetric.NewRC5("hello there i hate it here", 12))
	ciphersImplemented = append(ciphersImplemented, asymmetric.NewRSA())
	// Loop that encodes the messasge and then decodes it
	for _, cipher := range ciphersImplemented {
		encoded := cipher.Encode(Msg)
		fmt.Printf("Current cipher: %v\nEncoded message: %v\nDecoded message: %v\n\n", cipher.GetType(), encoded, cipher.Decode(encoded))
	}
	// The digital signature part
	testSign := signature.NewSignature(sha1.New())
	digiSign, digest := testSign.Sign(Msg)
	fmt.Printf("Digital signature:\n%v\n\n", digiSign)
	if testSign.Verify(digiSign, digest, testSign.PbKey, testSign.Mod) {
		fmt.Println("Signature verified and valid")
	} else {
		fmt.Println("Signature verified and invalid")
	}

}
