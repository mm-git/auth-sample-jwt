package main

import (
	"auth-sample-jwt/src/server"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	stopRequest := make(chan struct{})
	go signalHook(stopRequest)
	go server.Run(stopRequest)
	server.Wait()
}

func signalHook(stopRequest chan struct{}) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT)

	for {
		switch <-signalChan {
		case syscall.SIGINT:
			fmt.Println("SIGINT received. Application will stop after the next message received.")
			close(stopRequest)
			return
		default:
		}
	}
}
