package main

import (
	"crypto/sha256"
	"cs-labs/src/passwords"
	"cs-labs/src/server"
	"fmt"
	"net/http"
)

var (
	DataBase = passwords.NewUserDB(sha256.New())
)

func init() {
	fmt.Printf("This is a sample driver program to Show the results of the CS laboratory works\n")
}

func main() {
	DataBase := passwords.NewUserDB(sha256.New())
	DataBase.RegisterUser("Admin", "admin", "admin")
	DataBase.RegisterUser("Andrei", "1234", "user")
	fmt.Println(DataBase.Authenticate("Andrei", "1234"))
	http.HandleFunc("/atbash",
		server.Authenticate(
			server.Authorise(
				server.AtbashHandler,
				DataBase,
				"admin",
			),
			DataBase,
		),
	)
	http.ListenAndServe("localhost:8080", nil)
}
