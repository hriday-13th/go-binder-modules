// Client for TCP/IP protocol using net package
package protocols

import (
	"net"
)

func TCPClient() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Enter your message here
	message := "Hi! from client"
	conn.Write([]byte(message))
}