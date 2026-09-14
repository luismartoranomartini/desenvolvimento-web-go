package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ls, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Println("erro ao escutar", err)
		return
	}
	defer ls.Close()

	fmt.Println("conexão na porta 8080")

	for {
		var buff [1024]byte
		conn, err := ls.Accept()
		if err != nil {
			log.Println("erro ao aceitar a conexão:", err)
			conn.Close()
			continue
		}

		n, err := conn.Read(buff[:])
		if err != nil {
			log.Println("erro ao ler:", err)
			conn.Close()
			continue
		}
		fmt.Println("Bytes recebidos:", n)
		fmt.Println("Dados recebidos:", string(buff[:n]))

		msg := []byte("uma mensagem qualquer para ler")
		w, err := conn.Write(msg)
		if err != nil {
			log.Println("erro ao escrever:", err)
			conn.Close()
			continue
		}
		if w < len(msg) {
			log.Println("nem todos os bytes foram entregues:", w)
			conn.Close()
			continue
		}
		fmt.Println("Bytes enviados:", w)
		conn.Close()
	}
}
