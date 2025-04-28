package main

import (
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
		t.Errorf("Expected response to containe 'POST route matched', but got %s", response)
	}
}
