package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

const (
	serverAddr = "localhost:7777"
)

func main() {
	log.Println("Start web server")
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		webServer := SetupWebApp()
		log.Fatal(webServer.Listen(serverAddr))
	}()

	<-exit
	log.Println("Stopping app...")
}
