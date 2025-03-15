package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

type Response struct {
	body string
	code int
}

type Handler func(method, body string) Response

var routes map[string]Handler

func Start(port int) {
	p := ":" + strconv.Itoa(port)
	server, err := net.Listen("tcp", p)
	if err != nil {
		log.Println("err initiating server... " + err.Error())
	}

	for {
		socket, err := server.Accept()
		if err != nil {
			log.Println("err accepting socket")
		}

		go HandleClient(socket)
	}
}

func AddRoute(route string, handler Handler) {
	routes[route] = handler
}

func HandleClient(socket net.Conn) {
	defer socket.Close()

	rdr := bufio.NewReader(socket)
	requestLine, err := rdr.ReadString('\n')
	if err != nil {
		log.Println("err reading from socket... " + err.Error())
	}

	fmt.Println("Incoming request: " + requestLine)

	headers := make([]string, 0)
	var contentLength int = 0

	for {
		line, err := rdr.ReadString('\n')
		if err != nil {
			log.Println("err reading headers... " + err.Error())
			return
		}

		if strings.HasPrefix(line, "Content-Length") {
			parts := strings.Split(line, ":")
			lengthStr := strings.TrimSpace(parts[1])
			contentLength, err = strconv.Atoi(lengthStr)
		}

		line = strings.TrimSpace(line)

		if line == "" {
			break
		}

		headers = append(headers, line)
	}

	body := ""

	if contentLength > 0 {
		bodyBuffer := make([]byte, contentLength)
		_, err := io.ReadFull(rdr, bodyBuffer)
		if err != nil {
			log.Println("err reading body... " + err.Error())
			return
		}

		body = string(bodyBuffer)
	}

	requestParts := strings.Split(requestLine, " ")
	method := requestParts[0]
	endPoint := requestParts[1]

	handler := routes[endPoint]
	if handler != nil {
		response := handler(method, string(body))
		if method == "POST" {
			SendResponse(socket, response.body, "Created", response.code)
		} else if method == "GET" {
			SendResponse(socket, response.body, "OK", response.code)
		} else {
			SendResponse(socket, response.body, "Not Found", response.code)
		}
	} else {
		SendResponse(socket, "Not Found", "Not Found", 404)
	}
}

func SendResponse(socket net.Conn, body, message string, code int) {
	fmt.Fprintf(socket, "HTTP/1.1 %d %s\r\n", code, message)
	fmt.Fprintf(socket, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(socket, "Content-Type: text/plain\r\n")
	fmt.Fprintf(socket, "\r\n")
	fmt.Fprintf(socket, "%s", body)
}

func main() {
	routes = make(map[string]Handler)

	AddRoute("/home", func(method, body string) Response {
		if method == "GET" {
			return Response{
				body: "Hello World!",
				code: 200,
			}
		} else if method == "POST" {
			return Response{
				body: "You posted: " + body,
				code: 201,
			}
		}

		return Response{
			body: "Method Not Allowed",
			code: 406,
		}
	})

	Start(1069)
}
