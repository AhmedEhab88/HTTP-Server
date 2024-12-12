package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp4", ":8080")
	fmt.Println("Listening on 8080...")
	if err != nil {
		fmt.Printf("Error occured while listening. %v", err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error occured when Accepting. %v", err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the incoming data
	fmt.Printf("Received: %s", buf)
}
