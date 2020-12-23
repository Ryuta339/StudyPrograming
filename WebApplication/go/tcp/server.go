package main

import (
	"fmt"
	"io"
	"net"
)

func waitClient(listener net.Listener) {
	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}

	go goEcho(conn)
	waitClient(listener)
}

func goEcho(conn net.Conn) {
	defer conn.Close()
	echo(conn)
}

func echo(conn net.Conn) {
	var buf = make([]byte, 1024)

	n, err := conn.Read(buf)
	if err != nil {
		if err == io.EOF {
			return
		} else {
			panic(err)
		}
	}

	fmt.Printf("Client> %s \n", buf)
	n, err = conn.Write(buf[:n])
	if err != nil {
		panic(err)
	}

	echo(conn)
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8001")

	if err != nil {
		panic(err)
	}

	fmt.Println("Server running at localhost:8001")
	waitClient(listener)
}
