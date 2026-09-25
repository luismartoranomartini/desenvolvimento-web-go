package main

import (
	"io"
	"log"
	"net"
)

func handlerConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)

	for {
		nRead, errRead := conn.Read(buffer)

		if nRead > 0 {
			log.Printf("cliente enviou: %s\n", string(buffer[:nRead]))

			nWrite, errWrite := conn.Write(buffer[:nRead])
			if errWrite != nil {
				log.Println("erro ao enviar resposta:", errWrite)
				return
			}
			if nWrite != nRead {
				log.Printf("resposta parcial: enviados %d de %d bytes", nWrite, nRead)
				return
			}
		}
		if errRead != nil && errRead != io.EOF {
			log.Println("erro ao receber dados:", errRead)
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("erro ao estabeceler a conexão:", err)
	}
	defer listener.Close()

	log.Println("servidor escutando em localhost:8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("erro ao aceitar a conexão:", err)
			continue
		}
		go handlerConnection(conn)
	}
}
