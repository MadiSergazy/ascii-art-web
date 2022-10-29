package main

import (
	"Mado/asciiWeb/cmd/web"
	"log"
	"net/http"
)

func main() {
	// var ShowText string

	mux := http.NewServeMux()
	mux.HandleFunc("/", web.Home)

	// mux.HandleFunc("/snippet", web.ShowSnippet)
	// mux.HandleFunc("/snippet/create", web.CreateSnippet)

	log.Println("Запуск сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
