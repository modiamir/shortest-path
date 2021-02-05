package main

import (
	"log"
	"net/http"
	"shortest-path/http/handlers"
	"shortest-path/storage"
)

func main() {
	storage.SetDefaultStorage(storage.NewInMemoryStorage())

	http.Handle("/foo", handlers.ShortestPathHandler{})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
