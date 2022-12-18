package server

import (
	"cs-labs/src/ciphers/classical"
	"cs-labs/src/passwords"
	"fmt"
	"io"
	"net/http"
)

func Authenticate(next http.HandlerFunc, DB *passwords.UserDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if ok && DB.Authenticate(username, password) {
			next.ServeHTTP(w, r)
			return
		}
		http.Error(w, "None/Incorrect auth data provided", http.StatusUnauthorized)
	}
}

func Authorise(next http.HandlerFunc, DB *passwords.UserDB, allowedRoles ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userName, _, _ := r.BasicAuth()
		userRole, err := DB.GetRole(userName)
		if err == nil {
			for _, role := range allowedRoles {
				if role == userRole {
					next.ServeHTTP(w, r)
					return
				}
			}
		}
		http.Error(w, "Insufficient permissions", http.StatusForbidden)
	}
}
func AtbashHandler(w http.ResponseWriter, r *http.Request) {
	plainByte, err := io.ReadAll(r.Body)
	fmt.Println(string(plainByte))
	if err != nil {
		panic(err)
	}
	plaintext := string(plainByte)
	atbash := classical.NewAtbash("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	fmt.Fprintln(w, atbash.Encode(plaintext))
	return
}
