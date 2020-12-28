package main

import (
	"bufio"
	"fmt"
	"net"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

const (
	DOCUMENT_ROOT  = "/go/template"
	ERROR_DOCUMENT = "error_document"
	SERVER_NAME    = "0.0.0.0:8001"
)

func run(conn net.Conn) {
	var path string = ""
	var ext string = ""
	var host string = ""

	defer conn.Close()
	reader := bufio.NewReader(conn)
	line, err := reader.ReadString('\n')
	for err == nil {
		if line == "" || line == "\r\n" {
			break
		}
		if strings.HasPrefix(line, "GET") {
			path, _ = url.PathUnescape(strings.Split(line, " ")[1])
			tmp := strings.Split(path, ".")
			ext = tmp[len(tmp)-1]
		} else if strings.HasPrefix(line, "Host:") {
			host = line[:len("Host: ")]
		}
		line, err = reader.ReadString('\n')
	}
	if path == "" {
		return
	}

	if strings.HasSuffix(path, "/") {
		path += "index.html"
		ext = "html"
	}
	// Write Header
	writer := bufio.NewWriter(conn)
	realPath, err := filepath.Abs(DOCUMENT_ROOT + path)
	if err != nil {
		NewNotFoundResponse(ERROR_DOCUMENT).sendResponse(writer)
		return
	}
	fInfo, err := os.Stat(realPath)
	if err != nil {
		NewNotFoundResponse(ERROR_DOCUMENT).sendResponse(writer)
		return
	}

	if !strings.HasPrefix(realPath, DOCUMENT_ROOT) {
		NewNotFoundResponse(ERROR_DOCUMENT).sendResponse(writer)
		return
	} else if fInfo.IsDir() {
		var location string
		if host != "" {
			location = "http://" + host + path + "/"
		} else {
			location = "http://" + SERVER_NAME + path + "/"
		}
		NewMovePermanenltyResponse(location).sendResponse(writer)
		return
	}

	// Write Body
	fp, err := os.Open(realPath)
	if err != nil {
		NewNotFoundResponse(ERROR_DOCUMENT).sendResponse(writer)
		return
	}
	defer fp.Close()
	freader := bufio.NewReader(fp)
	NewOkResponse(freader, ext).sendResponse(writer)
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
