package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

const (
	CONN_HOSTNAME = "localhost"
	CONN_PORT     = "3333"
	CONN_HOST     = CONN_HOSTNAME + ":" + CONN_PORT
	CONN_TYPE     = "tcp"
)

func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buff := make([]byte, 1024)

	// Read the incoming connection into the buffer.
	_, err := conn.Read(buff)

	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	// print incoming data to the console
	fmt.Println(string(buff))

	// Send a response back to person contacting us.
	res := time.Now().Format("15:04:05") + "\n"
	conn.Write([]byte(res))

	// Close the connection when you're done with it.
	conn.Close()
}

func main() {
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST)

	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	// Close the listener when the application closes.
	defer l.Close()

	// print to console
	fmt.Println("tcp server listening on " + CONN_HOST)
	fmt.Println("hint: maybe try something like this...")
	fmt.Println("$ echo -n 'message' | nc " + CONN_HOSTNAME + " " + CONN_PORT)
	fmt.Println("")

	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()

		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}
