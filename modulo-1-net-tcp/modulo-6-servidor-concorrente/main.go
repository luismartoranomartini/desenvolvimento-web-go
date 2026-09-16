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
		n, errRead := conn.Read(buffer)

		_, errWrite := conn.Write(buffer[:n])
		if errWrite != nil {
			log.Println("erro ao escrever:", errWrite)
			return
		}
		if errRead != nil {
			if errRead == io.EOF {
				return
			}
			log.Println("erro ao ler", errRead)
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
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("erro ao aceitar a conexão", err)
			continue
		}
		fmt.Println("conexão foi aceita")
		go handlerConnection(conn)
	}

}
