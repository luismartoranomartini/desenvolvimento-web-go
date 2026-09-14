package main

import (
	"fmt"

	"log"

	"net"
)

func main() {

	ls, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Println("erro ao iniciar a escuta:", err)
		return
	}
	defer ls.Close()

	conn, err := ls.Accept()
	if err != nil {
		log.Println("erro ao aceitar a conexão")
		return
	}
	defer conn.Close()

	var buff [1024]byte

	fmt.Println("A conexão foi aceita!")

	n, err := conn.Read(buff[:])
	if err != nil {
		log.Println("erro ao ler:", err)
		return
	}

	fmt.Println("Bytes recebidos:", n)
	fmt.Println("Dados recebidos:", string(buff[:n]))

	mensagem := []byte("uma mensagem maior")

	w, err := conn.Write(mensagem)
	if err != nil {
		log.Println("erro ao enviar mensagem:", err)
		return
	}

	if w < len(mensagem) {
		log.Println("nem todos os bytes forma enviados:", w, "de,", len(mensagem))
		return
	}
	fmt.Println("Bytes enviados:", w)
}
