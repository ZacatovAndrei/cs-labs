# Public Key Cryptography

## Course: Cryptography & Security

## Author: Zacatov Andrei

----

### Theory

Public-key cryptography (also called asymmetric) is a field of cryptography that uses two different keys for encryption

- a keypair. Each pair consists of a private key that is kept secret, and a public key that is exchanged.
In such ciphers (called asymmetrical ciphers) everyone can encode the message with the provided public key, but only the
owner of the private key will be able to decrypt it.
Public-key cryptography can also be used for creating digital signatures, where the message would be encrypted with a
private key and everyone will be able to check the integrity with the help of the private key.

RSA is an example of such a system designed by Ron **R**ivest, Adi **S**hamir, and Leonard **A**dleman.
Rsa generates the keys based on two large prime numbers. The prime numbers are kept secret,encoded with the public key,
and decoded only with private keys.
RSA can also be used in the process of creating a digital signature.
The cryptographical strength of the cipher lies in the difficulty of factoring the product of large prime numbers.

RSA is a relatively slow algorithm. Because of this, it is not commonly used to directly encrypt user data. More often,
RSA is used to transmit shared keys for symmetric-key cryptography, which are then used for bulk encryption–decryption.

### Objectives

- Implement RSA encryption and decryption

### Implementation description

- Since the encryption process is a modular exponentiation problem the standard library of the Go programming language
  can be used.
  At the same time key generation of a decent length is also a problem, hence the standard library has been used again

- Key generation

```go
key, err := rsa.GenerateKey(rand.Reader, 2048)
if err != nil {
panic(err)
}
```

- Encryption implementation details:

1. since the cipher uses modular arithmetic it is impossible to decrypt the messages that are longer than the key
   without losing information.
2. The big.Int has an already implemented function of modular exponentiation, used in the `msg.EXP` part of the `return`
   statement.
   It takes 3 parameters that are base,exponent, and modulus
3. The output is the string of hexadecimal bytes

```go
func (R RSA) Encode(plain string) string {
if len(plain) > (2048 / 8) {
panic(errors.New("text too long to encode"))
}
msg := big.NewInt(0)
msg.SetBytes([]byte(plain))
return hex.EncodeToString(msg.Exp(msg, R.publicKey, R.modulus).Bytes())
}
```

- Decryption implementation details:
  the return value is cast to a string because of the constriction of the Cipher interface, defining return type of the
  function to be `string`

```go
func (R RSA) Decode(cipher string) string {
dCipher, err := hex.DecodeString(cipher)
if err != nil {
panic(err)
}
msg := big.NewInt(0)
msg.SetBytes(dCipher)
return string(msg.Exp(msg, R.privateKey, R.modulus).Bytes())
}
```

### Conclusions / Screenshots / Results

At the end of this laboratory work one has learned the principles and use cases of the asymmetric cryptography.
While implementing the cipher one could see why these sort of ciphers could not be used for encrypting whole messages
due to its complexity, and therefore is mostly used for digital signatures and transmitting keys for faster symmetrical
ciphers.
One could also use the implemented cipher to implement the digital signature algorithms later on in the laboratory work
chain
