package ciphers

type Cipher interface {
	GetType() string
	Encode(plainText string) string
	Decode(cipherText string) string
}
