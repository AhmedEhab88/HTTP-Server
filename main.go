package main

import (
	"fmt"
	"net"

	"golang.org/x/sys/windows"
)

func main() {
	socket, err := windows.Socket(windows.AF_INET, windows.SOCK_STREAM, windows.IPPROTO_TCP) // Create a new TCP Socket
	if err != nil {
		fmt.Printf("Error occured when creating socket %s", err)
		return
	}
	sockAddr := windows.SockaddrInet4{Port: 8080, Addr: [4]byte{127, 0, 0, 1}}

	err = windows.Bind(socket, &sockAddr) //Bind socket created to 127.0.0.1:8080

	if err != nil {
		fmt.Println("ERROR while binding socket to localhost")
		return
	}

	//Listen on socket
	err = windows.Listen(socket, 10) //Listening socket has backlog of 10 (Size of Accept Queue in Kernel)

	if err != nil {
		fmt.Println("ERROR while listening on socket")
		return
	}

	fmt.Println("Listening on 127.0.0.1:8080")

	//Accept incoming connection
	for {
		_, sa, err := windows.Accept(socket)

		if err != nil {
			fmt.Printf("Error while accepting new connection: %v", err)
			return
		}

		fmt.Printf("Address of connecting client: %s", sa)
	}

	// for {
	// 	conn, err := listener.Accept()
	// 	if err != nil {
	// 		fmt.Printf("Error occured when Accepting. %v", err)
	// 	}
	// 	go handleConnection(conn)
	// }
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
