package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func waitClient(listener net.Listener) {
	fmt.Println("Waiting a connection from a server...")
	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connect from a client.")

	go goEcho(conn)
	waitClient(listener)
}

func goEcho(conn net.Conn) {
	defer conn.Close()
	echo(conn)
}

func echo(conn net.Conn) {
	// var buf = make([]byte, 1024)

	/*
		_, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				return
			} else {
				panic(err)
			}
		}
	*/

	fp, err := os.Create("server_recv.txt")
	if err != nil {
		fmt.Println("File error.")
		return
	}
	defer fp.Close()

	output := bufio.NewWriter(fp)
	n, err := output.ReadFrom(conn)
	if err != nil {
		fmt.Println("Output error.")
		return
	}
	fmt.Println(n)
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8001")

	if err != nil {
		panic(err)
	}

	waitClient(listener)
}
