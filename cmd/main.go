package main

import (
	"mvcApiTest/pkg/api"
	"net/http"
)

func main() {
	http.HandleFunc("/api/users", api.GetUsers)
	http.HandleFunc("/api/users/register", api.CreateUsers)
	http.HandleFunc("/api/users/login", api.Login)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}
