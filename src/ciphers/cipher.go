package ciphers

type Cipher interface {
	GetType() string
	Encode(plain string) string
	Decode(cipher string) string
}
