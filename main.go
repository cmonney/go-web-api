package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", routeHandler)
	err := http.ListenAndServe("localhost:11111", nil)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func routeHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("Asset not found\n"))
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Running API V1\n"))
}
