package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	// inicia uma conexão
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	go func() {
		time.Sleep(5 * time.Second)

		err := listener.Close()
		if err != nil {
			log.Println("erro ao fechar o listener:", err)
		}
	}()
	_, err = listener.Accept()

	if errors.Is(err, net.ErrClosed) {
		fmt.Println("o listener foi fechado intencionalmente")
		return
	}
}
