package passwords

import (
	"encoding/hex"
	"fmt"
	"hash"
)

type UserDB struct {
	hashFunc    hash.Hash
	credentials map[string]string
}

func NewUserDB(hashFunc hash.Hash) *UserDB {
	return &UserDB{hashFunc: hashFunc, credentials: make(map[string]string)}
}

func (u *UserDB) RegisterUser(userName string, password string) {
	u.hashFunc.Write([]byte(password))
	hashedPassword := u.hashFunc.Sum(nil)
	u.credentials[userName] = hex.EncodeToString(hashedPassword)
	u.hashFunc.Reset()
}

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

func (u *UserDB) RemoveUser(username string) {
	delete(u.credentials, username)
}
