package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("erro ao escutar:", err)
	}
	defer conn.Close()

	msg := []byte("Olá, servidor!")
	_, errWrite := conn.Write(msg)
	if errWrite != nil {
		log.Println("erro ao escrever no cliente:", errWrite)
		return
	}
	buffer := make([]byte, 1024)
	nRead, errRead := io.ReadFull(conn, buffer)
	if err != nil {
		log.Printf("erro ao receber a resposta: %v, recebidos de %d de %bytes", errRead, nRead, len(buffer))
		return
	}
	fmt.Println(string(buffer[:nRead]))
}
