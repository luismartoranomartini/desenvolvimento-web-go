package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Println("erro ao abrir porta", err)
		return
	}
	defer listener.Close()

	var buffer [1024]byte

	fmt.Println("Servidor aguardando conexões na porta 8080...")

	conn, err := listener.Accept()
	if err != nil {
		log.Println("erro ao aceitar uma conexão:", err)
		return
	}
	defer conn.Close()

	fmt.Println("A conexão foi aceita!")

	n, err := conn.Read(buffer[:])
	if err != nil {
		log.Println("erro ao ler os dados", n)
		return
	}

	fmt.Println("Bytes recebidos: ", n)
	fmt.Println("Dados recebidos:", string(buffer[:n]))

}
