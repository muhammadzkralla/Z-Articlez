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

type Handler func(req Req, res Res)
type Middleware func(req Req, res Res, next func())

type Route struct {
	path    string
	handler Handler
}

type Res struct {
	socket net.Conn
	status int
}

type Req struct {
	method string
	path   string
	body   string
	params map[string]string
}

func (res *Res) send(data string) {
	sendResponse(res.socket, data, res.status)
}

var getRoutes []Route
var postRoutes []Route
var middlewares []Middleware

func Start(port int) {
	p := fmt.Sprintf(":%d", port)
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
	getRoutes = append(getRoutes, Route{endpoint, applyMiddleware(handler)})
}

func Post(endpoint string, handler Handler) {
	postRoutes = append(postRoutes, Route{endpoint, applyMiddleware(handler)})
}

func Use(m Middleware) {
	middlewares = append(middlewares, m)
}

func applyMiddleware(finalHandler Handler) Handler {
	return func(req Req, res Res) {
		i := 0

		var next func()
		next = func() {
			if i < len(middlewares) {
				middleware := middlewares[i]
				i++
				middleware(req, res, next)
			} else {
				finalHandler(req, res)
			}
		}

		next()
	}
}

func matchRoute(requestPath string, routes []Route) (Handler, map[string]string) {
	for _, route := range routes {
		params := make(map[string]string)
		routeParts := strings.Split(route.path, "/")
		requestParts := strings.Split(requestPath, "/")

		if len(routeParts) != len(requestParts) {
			continue
		}

		match := true
		for i := range routeParts {
			if strings.HasPrefix(routeParts[i], ":") {
				paramName := routeParts[i][1:]
				params[paramName] = requestParts[i]
			} else if routeParts[i] != requestParts[i] {
				match = false
				break
			}
		}

		if match {
			return route.handler, params
		}
	}

	return nil, nil
}

func handleClient(socket net.Conn) {
	defer socket.Close()

	rdr := bufio.NewReader(socket)
	requestLine, err := rdr.ReadString('\n')
	if err != nil {
		log.Println("err reading from socket... " + err.Error())
	}

	fmt.Println("Incoming request: " + requestLine)

	_, contentLength := extractHeaders(rdr)

	body := extractBody(rdr, contentLength)

	requestParts := strings.Split(requestLine, " ")
	method := requestParts[0]
	endPoint := requestParts[1]

	var handler Handler
	var params map[string]string

	if method == "GET" {
		handler, params = matchRoute(endPoint, getRoutes)
	} else if method == "POST" {
		handler, params = matchRoute(endPoint, postRoutes)
	}

	if handler != nil {
		req := Req{
			method: method,
			path:   endPoint,
			body:   body,
			params: params,
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

func extractHeaders(rdr *bufio.Reader) ([]string, int) {
	headers := make([]string, 0)
	var contentLength int = 0

	for {
		line, err := rdr.ReadString('\n')
		if err != nil {
			log.Println("err reading headers... " + err.Error())
			return nil, 0
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

	return headers, contentLength
}

func extractBody(rdr *bufio.Reader, contentLength int) string {

	body := ""

	if contentLength > 0 {
		bodyBuffer := make([]byte, contentLength)
		_, err := io.ReadFull(rdr, bodyBuffer)
		if err != nil {
			log.Println("err reading body... " + err.Error())
			return ""
		}

		body = string(bodyBuffer)
	}

	return body
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
