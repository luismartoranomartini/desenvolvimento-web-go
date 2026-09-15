package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Println("erro ao escutar:", err)
		return
	}
	defer listener.Close()

	for {

		conn, err := listener.Accept()
		if err != nil {
			log.Println("erro ao ler da conexão:", err)
			continue
		}
		go func(conn net.Conn) {
			defer conn.Close()
			var buff [1024]byte
			n, err := conn.Read(buff[:])
			if err != nil {
				log.Println("erro ao ler:", err)
				return
			}
			fmt.Println("Bytes recebidos:", n)
			fmt.Println("Dados recebidos", string(buff[:n]))

			msg := []byte("uma msg qualquer")
			w, err := conn.Write(msg)
			if err != nil {
				log.Println("erro ao escrever a mensagem:", err)
				return

			}
			if w < len(msg) {
				log.Println("nem todos os bytes foram enviados")
				return
			}
			fmt.Println("Bytes enviados:", w)

		}(conn)
	}
}
