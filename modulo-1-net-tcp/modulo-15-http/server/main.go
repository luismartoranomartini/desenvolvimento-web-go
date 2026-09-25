package main

import (
	"bufio"
	"log"
	"net"
	"strings"
)

func handlerConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	header := make(map[string]string)

	requestLine, err := reader.ReadString('\n')
	if err != nil {
		log.Println("erro ao ler a linha de requisição:", err)
		return
	}
	log.Printf("requisição: %s", strings.TrimSpace(requestLine))

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Println("erro ao ler os cabeçalhos:", err)
			return
		}
		line = strings.TrimRight(line, "\r\n")

		if line == "" {
			break
		}
		name, value, found := strings.Cut(line, ":")
		if !found {
			continue
		}
		name = strings.ToLower(strings.TrimSpace(name))
		value = strings.TrimSpace(value)
		header[name] = value
	}
	log.Println("Host:", header["host"])
	log.Println("Connection:", header["connection"])
	log.Println("Content-Length", header["content-length"])

	response := []byte(
		"HTTP/1.1 200 Ok\r\n" +
			"Content-Length: 3\r\n" +
			"Connection: close\r\n" +
			"\r\n" +
			"Ok\n",
	)
	_, err = conn.Write(response)
	if err != nil {
		log.Println("erro ao enviar resposta:", err)
		return
	}
}

func main() {
	listerner, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("erro ao escutar a conexão na porta 8080")
	}
	defer listerner.Close()

	for {
		conn, err := listerner.Accept()
		if err != nil {
			log.Println("erro ao aceitar a conexão:", err)
			continue
		}
		go handlerConnection(conn)
	}
}
