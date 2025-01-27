package main

import (
	"cyclic-code/src/backend"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Println("App start")

	go func() {
		if err := backend.StartServer(); err != nil {
			log.Fatalf("Could not start server: %v", err)
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	log.Println("App down")
}
