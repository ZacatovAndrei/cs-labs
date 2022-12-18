package passwords

import (
	"encoding/hex"
	"errors"
	"fmt"
	"hash"
)

type UserDB struct {
	hashFunc    hash.Hash
	credentials map[string]User
}

type User struct {
	HashedPass string
	Role       string
}

func NewUserDB(hashFunc hash.Hash) *UserDB {
	return &UserDB{hashFunc: hashFunc, credentials: make(map[string]User)}
}

func (u *UserDB) RegisterUser(userName string, password string, role string) {
	u.hashFunc.Write([]byte(password))
	hashedPassword := u.hashFunc.Sum(nil)
	u.credentials[userName] = User{
		HashedPass: hex.EncodeToString(hashedPassword),
		Role:       role,
	}
	u.hashFunc.Reset()
}

func (u *UserDB) GetRole(userName string) (string, error) {
	user, ok := u.credentials[userName]
	if !ok {
		return "", errors.New("no such username")
	}
	return user.Role, nil

}

func (u *UserDB) Authenticate(username string, password string) bool {
	u.hashFunc.Write([]byte(password))
	hashedPassword := hex.EncodeToString(u.hashFunc.Sum(nil))
	u.hashFunc.Reset()
	user, ok := u.credentials[username]
	if !ok || hashedPassword != user.HashedPass {
		fmt.Println("Error: Incorrect username or password")
		return false
	}
	if hashedPassword == user.HashedPass {
		fmt.Println("User successfully authorised")
		return true
	}
	return false
}

func (u *UserDB) RemoveUser(username string) {
	delete(u.credentials, username)
}
