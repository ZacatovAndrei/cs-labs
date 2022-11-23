# Hash Functions. Digital Signature

## Course: Cryptography & Security

## Author:  Andrei Zacatov

----

### Theory
A digital signature is a cryptographical scheme for verifying the authenticity and integrity of some data, be it messages or documents.
A valid signature tells the recipient that the message was created by a known sender and was not altered on its way.
Digital signatures use asymmetric cryptography under the hood.

A hash function is a one-way function that can map arbitrarily-sized data to a fixed-size value

### Objectives

1. Familiarise with the notion of hash functions
2. Use a previously implemented asymmetric cipher and implement the digital signature generation process
   1. Take the message 
   2. get a hash digest of it
   3. encrypt with the cipher used

3. Perform digital signature verification
   1. decrypt the provided message
   2. compare the result with the original message digest


### Implementation description
* Generating a digital signature  
```go
func (s Signature) Sign(message string) (string, string) {
    // Getting the hash sum of the message
	msg := []byte(message)
	s.digest.Write(msg)
	hashSum := s.digest.Sum(nil)
	...
	// Encoging the provided digest with RSA encryption scheme (the signature itself)
	out := msgNum.Exp(msgNum, s.prKey, s.Mod)
	//return the signature  and the digest for future use
	return out.Text(16), hex.EncodeToString(hashSum)
}
```

   * Some things to mention about the code
     1. The `digest` field of the Signature class is an another class, implementing a hash.Hash interface. 
     2. This interface provides a `Write` method, that is used to send data for hashing.
     3. The `Sum` Method that appends the sum to the argument, hence the use of `nil`
     4. The encoding resembles the one in the RSA class, but does not use it directly.That is mostly due to the  method signature definition in the `Cipher` interface.

* Digital signature verification process.  
The process of verifying a digital signature with the provided RSA encoding method requires the following parametres:
  * The signature itself
  * The message digest
  * Public key and Modulus pair obtained after generating an RSA key
 ```go

func (s Signature) Verify(message string, original string, publicKey *big.Int, modulus *big.Int) bool {
   // "Decryption" with public key
   res := msgNum.Exp(msgNum, publicKey, modulus)
   // if the digests coincide then signature is valid
   if res.Cmp(origMsg) != 0 {
   return false
   }
   return true
}
```
### Conclusions / Screenshots / Results

At the end of this laboratory work the process of creating and verifying the digital signatures had been studied.
The concept of a hash function has been studied, and we have realised that the combination of this family of functions and the previously studied asymmetrical cryptography can be used together to provide integrity and authenticity
