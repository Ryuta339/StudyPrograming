FROM golang:latest
WORKDIR /go
ADD . /go
RUN mkdir -p template
COPY template/* template/
RUN mkdir -p error_document
COPY error_document/* error_document
RUN go build -o server server.go send_response.go
CMD ["./server"]
