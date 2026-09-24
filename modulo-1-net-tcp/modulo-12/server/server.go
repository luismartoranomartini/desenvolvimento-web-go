package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)

func handlerConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	fmt.Println("o cliente conectado")

	for {
		nRead, errRead := conn.Read(buffer)

		if nRead > 0 {
			dadosRecebidos := buffer[:nRead]
			fmt.Printf("dados recebidos: %q", dadosRecebidos)
		}
		if errRead == nil {
			continue
		}
		if errors.Is(errRead, io.EOF) {
			fmt.Println("cliente encerrou a conexão")
			return
		}
		fmt.Println("erro inesperado na conexão:", errRead)
		return
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()

		if errors.Is(err, net.ErrClosed) {
			fmt.Println("listener encerrado")
			return
		}
		if err != nil {
			log.Println("erro ao aceitar a conexão:", err)
			continue
		}
		go handlerConnection(conn)
	}
}
