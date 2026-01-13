// Server for TCP/IP protocol using net package
package protocols

import (
	"fmt"
	"net"
)

func TCPServer() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080.....") // TBD: Parameterize it

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		goHandleConnection(conn)
	}
}

func goHandleConnection(conn net.Conn) {
	fmt.Println("Message Received!!")
}