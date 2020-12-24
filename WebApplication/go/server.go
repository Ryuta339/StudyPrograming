package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

const (
	DOCUMENT_ROOT = "template"
)

func writeLine(w *bufio.Writer, str string) error {
	_, err := w.WriteString(str + "\n")
	if err != nil {
		return err
	}
	err = w.Flush()
	return err
}

func run(conn net.Conn) {
	var path string
	var ext string

	defer conn.Close()
	reader := bufio.NewReader(conn)
	line, err := reader.ReadString('\n')
	for err == nil {
		if line == "" || line == "\r\n" {
			break
		}
		if strings.HasPrefix(line, "GET") {
			path = strings.Split(line, " ")[1]
			tmp := strings.Split(path, ".")
			ext = tmp[len(tmp)-1]
		}
		line, err = reader.ReadString('\n')
	}
	fmt.Println(path, ext)
	// Write Header
	writer := bufio.NewWriter(conn)
	writeLine(writer, "HTTP/1.1 200 OK")
	writeLine(writer, "Date: "+getDateStringUtc())
	writeLine(writer, "Server: Modoki/0.1")
	writeLine(writer, "Connection: close")
	writeLine(writer, "Content-Type: "+getContentType(ext))
	writeLine(writer, "")

	// Write Body
	fp, err := os.Open(DOCUMENT_ROOT + path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}
	defer fp.Close()
	writer.ReadFrom(fp)
	writer.Flush()
}

func getDateStringUtc() string {
	return time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05") + " GMT"
}

var contentTypeMap map[string]string = map[string]string{
	"html": "text/html",
	"htm":  "text/htm",
	"txt":  "text/plain",
	"css":  "text/css",
	"jpeg": "image/jpeg",
	"jpg":  "image/jpg",
	"png":  "image/png",
	"gif":  "image/gif",
}

func getContentType(ext string) (ret string) {
	ret = contentTypeMap[ext]
	if ret == "" {
		ret = "application/octet-stream"
	}
	return
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8001")
	if err != nil {
		panic(err)
	}

	for {
		fmt.Println("Server running at 0.0.0.0:8001")
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		run(conn)
	}
}
