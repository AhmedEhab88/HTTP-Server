package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println("Error while listening on port 8080")
		return
	}

	defer listener.Close()

	fmt.Println("Listening on 127.0.0.1:8080...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error occured when Accepting. %v", err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(connection net.Conn) {
	defer connection.Close()

	buf := make([]byte, 1024)
	_, err := connection.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := string(buf)
	request := strings.Split(data, "\r\n")
	requestSplitted := strings.Split(request[0], " ")

	result := fmt.Sprintf("HTTP/1.1 200 OK\r\n\r\nRequested path: %s\r\n", requestSplitted[1])

	_, err = connection.Write([]byte(result))

	if err != nil {
		fmt.Printf("Error while writing to connection: %s", err)
	}

}
