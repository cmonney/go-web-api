package main

import (
	"cmonney/go-web-api/handlers"
	"fmt"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/users/", handlers.UsersRouter)
	http.HandleFunc("/users", handlers.UsersRouter)
	http.HandleFunc("/", handlers.RouteHandler)
	err := http.ListenAndServe("localhost:11111", nil)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
