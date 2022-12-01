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

3. Implement secure password storage
   1. Create a database of "users". In this case it is an in-memory Hash map of form `map[string]string`, where the keys are logins and values are hash digests of passwords in a hex string format
   2. Implement ways of adding users and authenticating based on the username and password.

4. Perform digital signature verification
   1. decrypt the provided message
   2. compare the result with the original message digest

### Digital signature implementation description

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

   func (s Signature) Verify(message string, original string, publicKey *big.Int, modulus*big.Int) bool {
      // "Decryption" with public key
      res := msgNum.Exp(msgNum, publicKey, modulus)
      // if the digests coincide then signature is valid
      if res.Cmp(origMsg) != 0 {
      return false
      }
      return true
   }

   ```

### Password DB storage implementation

1. The Database is, as mentioned before, just a map\<string,string> with a hash method attached to it. The hashing method can be provided in the constructor.For the sake of security SHA-256 is used in the example.

   ```go
   //here one can see the structure of the DB, as well as the constructor
   type UserDB struct {
   hashFunc    hash.Hash
   credentials map[string]string
   }

   func NewUserDB(hashFunc hash.Hash) *UserDB {
   return &UserDB{
      hashFunc: hashFunc, 
      credentials: make(map[string]string)
      }
   }

   // the Constructor call in the provided example uses SHA-256.New() as the hash-function
      DataBase := passwords.NewUserDB(sha256.New())

   ```

2. This Database so far supports the following methods
   1. The RegisterUser method that would add the new (login, \<hash-set of password> to the database.

      ```go
      func (u *UserDB) RegisterUser(userName string, password string) {
         u.hashFunc.Write([]byte(password))
         hashedPassword := u.hashFunc.Sum(nil)
         u.credentials[userName] = hex.EncodeToString(hashedPassword)
         u.hashFunc.Reset()
      }
      ```

   2. The Authenticate method that returns a boolean `true` if the authentication was successful and `false` otherwise
         ```go
          func (u *UserDB) Authenticate(username string, password string) bool {
             u.hashFunc.Write([]byte(password))
             hashedPassword := hex.EncodeToString(u.hashFunc.Sum(nil))
             u.hashFunc.Reset()
             originalPassword, ok := u.credentials[username]
             if !ok || hashedPassword != originalPassword {
                fmt.Println("Error: Incorrect username or password")
                return false
             }
             if hashedPassword == originalPassword {
                fmt.Println("User successfully authenticated")
                return true
             }
             return false
          }
         ```
   3. The DeleteUser method that "deletes" a user from the database.The intention is to handle the whole verification part on the "frontend" side when user has already been authenticated and such, hence only the username will be needed.  
   For now, it mostly just deletes the entry from the map.
        ```go
            func (u *UserDB) RemoveUser(username string) {
            delete(u.credentials, username)
            }
        ```
   

3. Technically it would need a `Delete` method, however since it's an in-memory example there is no use for it in this context.`

### Conclusions / Screenshots / Results

At the end of this laboratory work the process of creating and verifying the digital signatures had been studied.
The concept of a hash function has been studied, and we have realised that the combination of this family of functions and the previously studied asymmetrical cryptography can be used together to provide integrity and authenticity
