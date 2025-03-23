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

type Res struct {
	socket net.Conn
	status int
}

type Req struct {
	method string
	path   string
	body   string
}

func (res *Res) send(data string) {
	sendResponse(res.socket, data, res.status)
}

type Handler func(req Req, res Res)

var getRoutes map[string]Handler
var postRoutes map[string]Handler

func init() {
	getRoutes = make(map[string]Handler)
	postRoutes = make(map[string]Handler)
}

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

		go handleClient(socket)
	}
}

func Get(endpoint string, handler Handler) {
	getRoutes[endpoint] = handler
}

func Post(endpoint string, handler Handler) {
	postRoutes[endpoint] = handler
}

func handleClient(socket net.Conn) {
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
	var handler Handler

	if method == "GET" {
		handler = getRoutes[endPoint]
	} else if method == "POST" {
		handler = postRoutes[endPoint]
	}

	if handler != nil {
		req := Req{
			method: method,
			path:   endPoint,
			body:   body,
		}
		res := Res{
			socket: socket,
			status: 200,
		}

		handler(req, res)
	} else {
		sendResponse(socket, "Not Found", 404)
	}
}

func sendResponse(socket net.Conn, body string, code int) {
	statusMessage := getHTTPStatusMessage(code)
	fmt.Fprintf(socket, "HTTP/1.1 %d %s\r\n", code, statusMessage)
	fmt.Fprintf(socket, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(socket, "Content-Type: text/plain\r\n")
	fmt.Fprintf(socket, "\r\n")
	fmt.Fprintf(socket, "%s", body)
}

func getHTTPStatusMessage(code int) string {
	statusMessages := map[int]string{
		200: "OK",
		201: "Created",
		400: "Bad Request",
		404: "Not Found",
		500: "Internal Server Error",
	}
	if msg, exists := statusMessages[code]; exists {
		return msg
	}
	return "Unknown Status"
}

func main() {
	Get("/home", func(req Req, res Res) {
		res.send("Hello, World!")
	})

	Post("/home", func(req Req, res Res) {
		reqBody := req.body
		respone := "You sent: " + reqBody
		res.send(respone)
	})

	Start(1069)
}
