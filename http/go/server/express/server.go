package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

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
