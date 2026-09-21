package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func hanclerConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		err := conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		if err != nil {
			log.Println("erro ao configurar o prazo de leitura:", err)
		}
		rRead, errRead := conn.Read(buffer)

		if rRead > 0 {
			fmt.Printf("Bytes recebidos: %q\n", rRead)
			fmt.Printf("Dados recebidis: %q\n", string(buffer[:rRead]))

			wWrite, errWrite := conn.Write(buffer[:rRead])
			if errWrite != nil {
				log.Println("erro ao escrever no servidor:", errWrite)
				return
			}
			fmt.Printf("Bytes enviados: %d\n", wWrite)

		}
		if errRead != nil {
			if errRead == io.EOF {
				log.Println("o cliente encerrou a conexão")
				return
			}
			if netErr, ok := errRead.(net.Error); ok && netErr.Timeout() {
				log.Println("tempo limite de leitura atingido")
				return
			}
			log.Println("erro ao ler do cliente:", errRead)
			return
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("erro ao aceitar a conexão:", err)
			continue
		}
		go hanclerConnection(conn)
	}

}
