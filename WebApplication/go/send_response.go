package main

import (
	"bufio"
	"os"
	"time"
)

func writeLine(w *bufio.Writer, str string) error {
	_, err := w.WriteString(str + "\n")
	if err != nil {
		return err
	}
	err = w.Flush()
	return err
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

func sendOkResponse(output *bufio.Writer, input *bufio.Reader, ext string) {

	writeLine(output, "HTTP/1.1 200 OK")
	writeLine(output, "Date: "+getDateStringUtc())
	writeLine(output, "Server: Modoki/0.1")
	writeLine(output, "Connection: close")
	writeLine(output, "Content-Type: "+getContentType(ext))
	writeLine(output, "")

	output.ReadFrom(input)
}

func sendMovePermanentlyResponse(output *bufio.Writer, location string) {
	writeLine(output, "HTTP/1.1 301 Moved Permanently")
	writeLine(output, "Date: "+getDateStringUtc())
	writeLine(output, "Server: Modoki/0.2")
	writeLine(output, "Location: "+location)
	writeLine(output, "Connection: close")
	writeLine(output, "")
}

func sendNotFoundResponse(output *bufio.Writer, errorDocumentRoot string) {
	writeLine(output, "HTTP/1.1 404 Not Found")
	writeLine(output, "Date: "+getDateStringUtc())
	writeLine(output, "Server: Modoki/0.1")
	writeLine(output, "Connection: close")
	writeLine(output, "Content-type: text/html")
	writeLine(output, "")

	// Write Body
	fp, err := os.Open(errorDocumentRoot + "/404.html")
	if err != nil {
		return
	}
	defer fp.Close()
	reader := bufio.NewReader(fp)
	output.ReadFrom(reader)
}
