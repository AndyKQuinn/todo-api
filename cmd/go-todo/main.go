package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AndyGuice/todo-api/api/server"
	"github.com/AndyGuice/todo-api/config"
)

func main() {
	config.Initialize()
	r := server.Router()
	fmt.Println("Starting server on the port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
