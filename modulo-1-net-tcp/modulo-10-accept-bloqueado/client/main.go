package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// inicia uma conexão
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Println("erro ao estabelecer a conexão:", err)
		return
	}
	defer conn.Close()

	mensagem := []byte("Olá, servidor!")
	_, errWrite := conn.Write(mensagem)
	if err != nil {
		log.Println("erro ao escrever no cliente:", errWrite)
		return
	}
	buffer := make([]byte, 1024)

	nRead, errRead := conn.Read(buffer)
	if errRead != nil {
		log.Println("erro ao receber:", errRead)
		return
	}
	fmt.Println(string(buffer[:nRead]))
}
