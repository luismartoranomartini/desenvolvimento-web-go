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
			_, errWrite := conn.Write(buffer[:r])
			if errWrite != nil {
				log.Println("erro ao escrever no servidor", errWrite)
				return
			}
		}
		if errRead != nil {
			if errRead == io.EOF {
				log.Println("o cliente encerrou a conexão")
				return
			}
			log.Println("erro ao ler o servidor:", errRead)
			return
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println("erro na escuta:", err)
		return
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("erro ao aceitar a conexão", err)
			continue
		}
		fmt.Println("Servidor conectado!")
		go handlerConnection(conn)
	}
}
