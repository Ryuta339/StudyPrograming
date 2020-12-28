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

type Response interface {
	sendResponse(output *bufio.Writer)
}

type OkResponse struct {
	input *bufio.Reader
	ext   string
}

func NewOkResponse(input *bufio.Reader, ext string) *OkResponse {
	return &OkResponse{
		input: input,
		ext:   ext,
	}
}

func (resp *OkResponse) sendResponse(output *bufio.Writer) {

	writeLine(output, "HTTP/1.1 200 OK")
	writeLine(output, "Date: "+getDateStringUtc())
	writeLine(output, "Server: Modoki/0.1")
	writeLine(output, "Connection: close")
	writeLine(output, "Content-Type: "+getContentType(resp.ext))
	writeLine(output, "")

	output.ReadFrom(resp.input)
}

type MovePermanentlyResponse struct {
	location string
}

func NewMovePermanenltyResponse(location string) *MovePermanentlyResponse {
	return &MovePermanentlyResponse{
		location: location,
	}
}

func (resp *MovePermanentlyResponse) sendResponse(output *bufio.Writer) {
	writeLine(output, "HTTP/1.1 301 Moved Permanently")
	writeLine(output, "Date: "+getDateStringUtc())
	writeLine(output, "Server: Modoki/0.2")
	writeLine(output, "Location: "+resp.location)
	writeLine(output, "Connection: close")
	writeLine(output, "")
}

type NotFoundResponse struct {
	errorDocumentRoot string
}

func NewNotFoundResponse(errorDocumentRoot string) *NotFoundResponse {
	return &NotFoundResponse{
		errorDocumentRoot: errorDocumentRoot,
	}
}

func (resp *NotFoundResponse) sendResponse(output *bufio.Writer) {
	writeLine(output, "HTTP/1.1 404 Not Found")
	writeLine(output, "Date: "+getDateStringUtc())
	writeLine(output, "Server: Modoki/0.1")
	writeLine(output, "Connection: close")
	writeLine(output, "Content-type: text/html")
	writeLine(output, "")

	// Write Body
	fp, err := os.Open(resp.errorDocumentRoot + "/404.html")
	if err != nil {
		return
	}
	defer fp.Close()
	reader := bufio.NewReader(fp)
	output.ReadFrom(reader)
}
