package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Println("erro ao iniciar o cliente:", err)
		return
	}

	msg := []byte("mensagem")
	_, errWrite := conn.Write(msg)
	if errWrite != nil {
		log.Println("erro ao escrever no cliente:", errWrite)
		return
	}
	buffer := make([]byte, 1024)
	rRead, errRead := conn.Read(buffer)
	if errRead != nil {
		log.Println("erro ao ler no cliente:", errRead)
		return
	}
	fmt.Println("Bytes enviados:", rRead)
	fmt.Println("Dados enviados:", string(rRead))

}
