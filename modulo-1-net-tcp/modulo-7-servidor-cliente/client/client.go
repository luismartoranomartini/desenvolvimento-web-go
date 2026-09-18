package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Println("erro ao estabelecer o conexão:", err)
		return
	}
	defer conn.Close()

	w, errWrite := conn.Write([]byte("Olá, servidor"))
	if errWrite != nil {
		log.Println("erro na escrita", errWrite)
		return
	}

	fmt.Printf("Bytes enviados: %d\n", w)
	buffer := make([]byte, 1024)

	r, errRead := conn.Read(buffer)
	if errRead != nil {
		log.Println("erro ao ler", errRead)
		return
	}

	fmt.Printf("Bytes recebidos: %d\n", r)
	fmt.Printf("Resposta do servidor: %q\n", string(buffer[:r]))
}
