package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Println("erro ao iniciar:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Servidor aguardando conexões na porta 8080...")

	conn, err := listener.Accept()
	if err != nil {
		log.Println("erro ao aceitar a conexão", err)
		return
	}
	defer conn.Close()
	fmt.Println("A conexão foi aceita!")

	mensagem := []byte("Olá, cliente")
	n, err := conn.Write(mensagem)
	if err != nil {
		log.Println("erro ao escrever:", err)
		return
	}
	if n < len(mensagem) {
		log.Println("a chamada não aceitou todos os bytes da mensagem.")
		return
	}
	fmt.Println("Bytes recebidos:", n)

}
