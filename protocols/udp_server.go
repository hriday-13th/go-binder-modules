// Server for UDP protocol using net package
package protocols

import (
	"fmt"
	"net"
)

func UDPServer() {
	addr, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		panic(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Server is listening on 8080....")

	buffer := make([]byte, 1024)

	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			panic(err)
		}

		message := string(buffer[:n])
		fmt.Println("Message received")

		response := fmt.Sprintf("Server received - %s\n", message)
		_, err = conn.WriteToUDP([]byte(response), clientAddr)
		if err != nil {
			panic(err)
		}
	}
}