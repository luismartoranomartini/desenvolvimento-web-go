package main

import (
	"errors"
	"io"
	"log"
	"net"
)

func writeAll(w io.Writer, dados []byte) error {
	for len(dados) > 0 {
		nWrite, errWrite := w.Write(dados)

		if errWrite != nil {
			return errWrite
		}
		if nWrite == 0 {
			return io.ErrShortWrite
		}
		dados = dados[nWrite:]
	}
	return nil
}

func handlerConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		nRead, errRead := conn.Read(buffer)

		if nRead > 0 {
			dadosRecebidos := buffer[:nRead]

			nWrite, errWrite := conn.Write(dadosRecebidos)

			if errWrite != nil {
				log.Printf("erro ao escrever na conexão: %v, escritos %d de %d bytes", errWrite, nWrite, len(dadosRecebidos))
				return
			}
			if nWrite != len(dadosRecebidos) {
				log.Printf("escrita incompleta: escritos %d de %d bytes", nWrite, len(dadosRecebidos))
				return
			}
		}
		if errRead == nil {
			continue
		}
		if errors.Is(errRead, io.EOF) {
			return
		}
		log.Println("erro ao ler da conexão:", errRead)
		return
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("erro ao escutar:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("erro ao aceitar a conexão:", err)
			continue
		}
		go handlerConnection(conn)
	}
}
