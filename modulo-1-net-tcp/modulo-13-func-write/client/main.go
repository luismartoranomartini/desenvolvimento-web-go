package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	msg := []byte("Olá, servidor!")
	nWrite, errWrite := conn.Write(msg)
	if nWrite != len(msg) {
		log.Println("erro enviar:", errWrite)
		return
	}
	buffer := make([]byte, 1024)
	totalLido := 0

	for totalLido < len(msg) {
		nRead, errRead := conn.Read(buffer[totalLido:])
		totalLido += nRead

		if totalLido == len(msg) {
			break
		}
		if errRead != nil {
			log.Println("erro ao receber:", errRead)
			return
		}
	}
	fmt.Println(string(buffer[:totalLido]))
}
