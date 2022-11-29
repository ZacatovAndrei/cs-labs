package main

import (
	"crypto/sha256"
	"cs-labs/src/passwords"
	"cs-labs/src/signature"
	"fmt"
)

const (
	Msg = "This is some Plaintext. Also, the classical ciphers loose all case-sensitive Data with classical ciphers."
)

func init() {
	fmt.Printf("This is a sample driver program to Show the results of the CS laboratory works\n")
	fmt.Printf("Author:Zacatov Andrei\nGroup:FAF-201\n")
	fmt.Printf("Test message is: %v\n\n", Msg)
}

func main() {
	/*
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
	*/
	// The digital signature part
	testSign := signature.NewSignature(sha256.New())
	digitalSignature, digest := testSign.Sign(Msg)
	fmt.Printf("Digital signature:\n%v\n\n", digitalSignature)
	if testSign.Verify(digitalSignature, digest, testSign.PbKey, testSign.Mod) {
		fmt.Println("Signature verified and valid")
	} else {
		fmt.Println("Signature verified and invalid")
	}
	DataBase := passwords.NewUserDB(sha256.New())
	DataBase.RegisterUser("admin", "password")
	DataBase.RegisterUser("Andrei", "FAF-201")

	fmt.Println("test Function to add user with user input (no input whitelisting")
	addUser(DataBase)
	fmt.Println("Test function to authenticate a user with user input (No input whitelisting)")
	if success := authenticate(DataBase); success {
		fmt.Println("Login successful")
	} else {
		fmt.Println("Login unsuccessful")
	}
}

func addUser(db *passwords.UserDB) {
	var (
		username = ""
		password = ""
	)
	fmt.Printf("username:")
	fmt.Scanln(&username)
	fmt.Printf("password:")
	fmt.Scanln(&password)
	db.RegisterUser(username, password)
}

func authenticate(db *passwords.UserDB) bool {
	var (
		username = ""
		password = ""
	)
	fmt.Printf("username:")
	fmt.Scanln(&username)
	fmt.Printf("password:")
	fmt.Scanln(&password)
	return db.Authenticate(username, password)
}
