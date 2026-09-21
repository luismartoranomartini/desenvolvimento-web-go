package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func handlerConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)

	for {
		r, errRead := conn.Read(buffer)
		if r > 0 {
			fmt.Println("Bytes recebidos:", r)
			fmt.Println("Dados recebidos:", string(buffer[:r]))

			_, errWrite := conn.Write(buffer[:r])
			if errWrite != nil {
				log.Println("erro ao escrever no servidor:", errWrite)
				return
			}
		}
		if errRead != nil {
			if errRead == io.EOF {
				log.Println("cliente encerrou a conexão")
				return
			}
			log.Println("erro ao ler no servidor:", errRead)
			return
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("erro ao escutar", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("erro ao aceitar a conexão:", err)
			continue
		}
		go handlerConnection(conn)
	}
}
