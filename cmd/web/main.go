package main

import (
	"encoding-restoration/internal/api"
	"log"
)

func main() {
	log.Println("Application started!")
	api.StartServer()
	log.Println("Application terminated!")
}
