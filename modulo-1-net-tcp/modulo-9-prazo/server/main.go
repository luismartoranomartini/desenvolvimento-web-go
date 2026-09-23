package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func handlerConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		err := conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		if err != nil {
			log.Println("erro ao configurar o deadline:", err)
			return
		}
		nRead, errRead := conn.Read(buffer)

		if nRead > 0 {
			fmt.Println("Bytes enviados:", nRead)
			fmt.Println("Dados enviados:", string(buffer[:nRead]))

			nWrite, errWrite := conn.Write(buffer[:nRead])
			if errWrite != nil {
				log.Println("erro ao escrever no servidor:", errWrite)
				return
			}
			fmt.Println("Bytes recebidos:", nWrite)
		}
		if errRead == nil {
			continue
		}

		var netErr net.Error

		if errors.Is(errRead, io.EOF) {
			log.Println("o cliente encerrou a conexão")
			return
		}
		if errors.As(errRead, &netErr) && netErr.Timeout() {
			log.Println("o tempo limite de leitura excedido")
			return
		}
		log.Println("erro ao ler no servidor:", errRead)
		return
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("erro ao escutar:", err)
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
