package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/zuzuqo/go-bookstore/pkg/routes"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Printf("Starting server at port 9000\n")
	log.Fatal(http.ListenAndServe(":9000", r))
}
