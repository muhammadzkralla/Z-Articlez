package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"strings"
	"testing"
	"time"
)

// Mock of net.Conn struct following the net.Conn interface specifications
type MockConn struct {
	data []byte
}

// All these functions are mocked to follow the net.Conn interface specifications
// for testing purposes only
func (m *MockConn) Read(p []byte) (n int, err error) {
	copy(p, m.data)
	return len(m.data), nil
}

func (m *MockConn) Write(p []byte) (n int, err error) {
	m.data = append(m.data, p...)
	return len(p), nil
}

func (m *MockConn) Close() error {
	return nil
}

func (m *MockConn) LocalAddr() net.Addr {
	return nil
}

func (m *MockConn) RemoteAddr() net.Addr {
	return nil
}

func (m *MockConn) SetDeadline(t time.Time) error {
	return nil
}

func (m *MockConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (m *MockConn) SetWriteDeadline(t time.Time) error {
	return nil
}

// Helper function to mock a request and return the response
func mockRequest(method, path, body string) string {
	conn := &MockConn{}
	req := fmt.Sprintf("%s %s HTTP/1.1\r\nContent-Length: %d\r\n\r\n%s", method, path, len(body), body)
	conn.data = []byte(req)

	// Call handleClient with the mocked connection
	handleClient(conn)

	return string(conn.data)
}

// Test GET route matching
func TestGetRouteMatching(t *testing.T) {
	// Mock a GET handler
	Get("/test", func(req Req, res Res) {
		res.send("GET route matched")
	})

	// Mock a GET request
	response := mockRequest("GET", "/test", "")

	if !strings.Contains(response, "GET route matched") {
		t.Errorf("Expected response to contain 'GET route matched', but got %s", response)
	}
}

// Test POST route matching
func TestPostRouteMatching(t *testing.T) {
	// Mock a POST handler
	Post("/test", func(req Req, res Res) {
		res.send("POST route matched")
	})

	// Mock a POST request
	response := mockRequest("POST", "/test", "")

	if !strings.Contains(response, "POST route matched") {
		t.Errorf("Expected response to contain 'POST route matched', but got %s", response)
	}
}

// Test middleware
func TestMiddleware(t *testing.T) {
	// Mock a middleware
	Use(func(req Req, res Res, next func()) {
		res.send("Middleware worked")
		next()
	})

	// Mock a handler
	Get("/test", func(req Req, res Res) {

	})

	response := mockRequest("GET", "/test", "")

	if !strings.Contains(response, "Middleware worked") {
		t.Errorf("Expected response to contain 'Middleware worked', but got %s", response)
	}
}

// Test dynamic routing
func TestDynamicRouting(t *testing.T) {
	// Mock a GET handler
	Get("/test/:postId/comment/:commentId", func(req Req, res Res) {
		postId := req.params["postId"]
		commentId := req.params["commentId"]
		res.send("Post ID: " + postId + ", Comment ID: " + commentId)
	})

	// Mock a POST handler
	Post("/test/:postId/comment/:commentId", func(req Req, res Res) {
		postId := req.params["postId"]
		commentId := req.params["commentId"]
		res.send("Post ID: " + postId + ", Comment ID: " + commentId)
	})

	getResponse := mockRequest("GET", "/test/123/comment/comment1", "")
	postResponse := mockRequest("POST", "/test/123/comment/comment1", "")

	if !strings.Contains(getResponse, "Post ID: 123, Comment ID: comment1") {
		t.Errorf("Expected response to contain 'Post ID:123, Comment ID:comment1', but got %s", getResponse)
	}

	if !strings.Contains(postResponse, "Post ID: 123, Comment ID: comment1") {
		t.Errorf("Expected response to contain 'Post ID:123, Comment ID:comment1', but got %s", postResponse)
	}
}

func TestSendResponse(t *testing.T) {
	conn := &MockConn{}
	res := Res{
		socket: conn,
		status: 200,
	}

	res.send("OK")

	expected := "HTTP/1.1 200 OK\r\nContent-Length: 2\r\nContent-Type: text/plain\r\n\r\nOK"
	if string(conn.data) != expected {
		t.Errorf("Expected response %s, but got %s", expected, string(conn.data))
	}
}

func TestExtractHeaders(t *testing.T) {
	headers := "Content-Length: 20\r\nHeader1: header1\r\nHeader2: header2\r\n\r\n"
	rdr := bufio.NewReader(bytes.NewBufferString(headers))

	extractedHeaders, extractedLen := extractHeaders(rdr)

	if extractedHeaders[0] != "Content-Length: 20" {
		t.Errorf("Expected header 'Content-Length: 20', but got %s", extractedHeaders[0])
	}

	if extractedHeaders[1] != "Header1: header1" {
		t.Errorf("Expected header 'Header1: header1', but got %s", extractedHeaders[1])
	}

	if extractedHeaders[2] != "Header2: header2" {
		t.Errorf("Expected header 'Header2: header2', but got %s", extractedHeaders[2])
	}

	if extractedLen != 20 {
		t.Errorf("Expected Content-Length 20, but got %d", extractedLen)
	}
}

func TestExtractBody(t *testing.T) {
	body := "Hello, world!"
	rdr := bufio.NewReader(bytes.NewBufferString(body))

	extractedBody := extractBody(rdr, 13)

	if extractedBody != "Hello, world!" {
		t.Errorf("Expected body to be 'Hello, world!', but got %s", extractedBody)
	}
}
