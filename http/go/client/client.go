package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

type OnComplete func(success any, failure string)

var baseUrl string

func InitClient(url string) {
	baseUrl = url
}

func connect() net.Conn {
	socket, err := net.Dial("tcp", baseUrl)
	if err != nil {
		log.Println("err creating socket... " + err.Error())
	}

	return socket
}

func Get(endpoint string, onComplete OnComplete) {
	socket := connect()
	defer socket.Close()

	requestLine := "GET " + endpoint + " HTTP/1.1\r\n"
	host := "Host: " + strings.Replace(baseUrl, "http://", "", -1) + "\r\n"
	userAgent := "User-Agent: zclient\r\n"
	accept := "Accept: */*\r\n\r\n"

	request := requestLine + host + userAgent + accept

	_, err := socket.Write([]byte(request))
	if err != nil {
		log.Println("err sending message... " + err.Error())
	}

	rdr := bufio.NewReader(socket)
	responseLine, err := rdr.ReadString('\n')
	responseParts := strings.Split(responseLine, " ")

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

	code := responseParts[1]
	statusCode, err := strconv.Atoi(code)

	if statusCode >= 200 && statusCode < 300 {
		onComplete(body, "")
	} else {
		onComplete("", "err "+body)
	}
}

func Post(endpoint string, body any, onComplete OnComplete) {
	socket := connect()
	defer socket.Close()

	jsonBody, err := json.Marshal(body)
	if err != nil {
		log.Println("err serializing body... " + err.Error())
		onComplete(nil, err.Error())
		return
	}

	bodyContentLength := len(jsonBody)

	requestLine := "POST " + endpoint + " HTTP/1.1\r\n"
	host := "Host: " + strings.Replace(baseUrl, "http://", "", -1) + "\r\n"
	userAgent := "User-Agent: zclient\r\n"
	accept := "Accept: */*\r\n"
	contentType := "Content-Type: text/plain\r\n"
	contentLengthHeader := "Content-Length: " + strconv.Itoa(bodyContentLength) + "\r\n\r\n"

	request := requestLine + host + userAgent + accept + contentType + contentLengthHeader + string(jsonBody)

	_, err = socket.Write([]byte(request))
	if err != nil {
		log.Println("err sending message... " + err.Error())
	}

	rdr := bufio.NewReader(socket)
	responseLine, err := rdr.ReadString('\n')
	responseParts := strings.Split(responseLine, " ")

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

	responseBody := ""

	if contentLength > 0 {
		bodyBuffer := make([]byte, contentLength)
		_, err := io.ReadFull(rdr, bodyBuffer)
		if err != nil {
			log.Println("err reading body... " + err.Error())
			return
		}

		responseBody = string(bodyBuffer)
	}

	code := responseParts[1]
	statusCode, err := strconv.Atoi(code)

	if statusCode >= 200 && statusCode < 300 {
		onComplete(responseBody, "")
	} else {
		onComplete("", "err "+responseBody)
	}
}

func main() {
	InitClient("localhost:1069")

	Get("/home", func(success any, failure string) {
		if success != nil {
			fmt.Println("Success: ", success)
		} else {
			fmt.Println("Failure: ", failure)
		}
	})

	// jsonBody := "\"{\"name\": \"John\", \"age\": 30}\""

	Post("/home", "jsonBody", func(success any, failure string) {
		if success != nil {
			fmt.Println("Success: ", success)
		} else {
			fmt.Println("Failure: ", failure)
		}
	})
}
