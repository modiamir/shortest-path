package main

import (
	"github.com/modiamir/shortest-path/http/handlers"
	"github.com/modiamir/shortest-path/storage"
	"log"
	"net/http"
)

func main() {
	storage.SetDefaultStorage(storage.NewInMemoryStorage())

	http.Handle("/foo", handlers.ShortestPathHandler{})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
