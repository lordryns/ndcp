package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

var clients []net.Conn

func main() {
	log.Println("Starting server...")
	listener, err := net.Listen("tcp", ":8118")
	if err != nil {
		log.Fatal("err: ", err.Error())
	}

	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("err: ", err.Error())
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	clients = append(clients, conn)
	reader := bufio.NewReader(conn)
	sender := conn.RemoteAddr().String()

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Println("err: ", err.Error())
		}

		response := fmt.Sprintf("message from (%v): %v", sender, message)
		for _, client := range clients {
			if client != conn {
				client.Write([]byte(response))
			}
		}

	}
}
