package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("erro ao estabelecer a conexão:", err)
	}
	defer conn.Close()

	msg := []byte("GET / HTTP/1.1\r\nHost: localhost:8080\r\nConnection: close\r\n\r\n")
	_, writeErr := conn.Write(msg)
	if writeErr != nil {
		log.Println("erro ao escrever do cliente:", writeErr)
		return
	}
	buffer := make([]byte, 1024)

	nRead, readErr := conn.Read(buffer)
	if readErr != nil {
		log.Printf("erro ao receber a resposta: %v, recebidos de %d de %bytes", readErr, nRead, len(buffer))
		return
	}
	fmt.Println(string(buffer[:nRead]))
}
