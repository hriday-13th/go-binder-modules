// Client for UDP protocol using net package
package protocols

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func UDPClient() {
	serverAddr, err := net.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Connected to the UDP server!!")

	scanner := bufio.NewScanner(os.Stdin)
	buffer := make([]byte, 1024)

	for {
		if !scanner.Scan() {
			break
		}

		message := strings.TrimSpace(scanner.Text())
		if message == "quit" {
			break
		}

		_, err := conn.Write([]byte(message))
		if err != nil {
			panic(err)
		}

		n, err := conn.Read(buffer)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Message recevied by Server - %s\n", string(buffer[:n]))
	}
	fmt.Println("Connection Closed.")
}