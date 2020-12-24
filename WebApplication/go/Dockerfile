FROM golang:latest
WORKDIR /go
ADD . /go
RUN mkdir -p template
COPY template/* template/
RUN go build -o server server.go
CMD ["./server"]
