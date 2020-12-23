package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func sendMessage(conn net.Conn) {
	fmt.Print("> ")

	stdin := bufio.NewScanner(os.Stdin)
	if stdin.Scan() == false {
		fmt.Println("Ciao ciao!")
		return
	}

	_, err := conn.Write([]byte(stdin.Text()))
	if err != nil {
		panic(err)
	}

	var resp = make([]byte, 4*1024)
	_, err = conn.Read(resp)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Server> %s \n", resp)
	sendMessage(conn)
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8001")

	if err != nil {
		panic(err)
	}

	defer conn.Close()
	sendMessage(conn)
}
