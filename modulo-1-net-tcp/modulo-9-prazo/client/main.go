package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Println("erro ao estabelecer a conexão")
		return
	}
	defer conn.Close()

	err = conn.SetDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		log.Println("erro ao configurar a o Deadline:", err)
		return
	}
	msg := []byte("Olá, servidor!")
	nWrite, errWrite := conn.Write(msg)
	if errWrite != nil {
		log.Println("erro ao escrever no cliente:", errWrite)
		return
	}
	log.Println("Bytes recebidos:", nWrite)

	buffer := make([]byte, 1024)
	nRead, errRead := conn.Read(buffer)
	if errRead != nil {
		var netErr net.Error
		if errors.As(errRead, &netErr) && netErr.Timeout() {
			log.Println("tempo de leitura excedido")
			return
		}
	}
	fmt.Println("Bytes recebidos:", nRead)
	fmt.Println("Dados recebidos:", string(buffer[:nRead]))
}
