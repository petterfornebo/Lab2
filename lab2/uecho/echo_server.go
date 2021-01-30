// Leave an empty line above this comment.
package main

import (
	"log"
	"net"
	"strings"
	"unicode"
)

type UDPServer struct {
	conn *net.UDPConn
	// TODO(student): Add fields if needed
}

// UDPServer implements the UDP Echo Server specification found at
// https://github.com/COURSE_TAG/assignments/tree/master/lab2/README.md#udp-echo-server

func (u *UDPServer) ServeUDP() {
	// TODO(student): Implement
	buf := make([]byte, 65535)
	defer u.conn.Close()

	for {
		n, raddr, err := u.conn.ReadFromUDP(buf)
		if err != nil {
			if socketIsClosed(err) {

			}
			log.Println("Reading packet...", err)
			continue
		}
		err = u.handleCommand(buf[:n], raddr)
		if err != nil {
			log.Println(err)
		}

	}
}
func NewUDPServer(addr string) (*UDPServer, error) {
	// TODO(student): Implement
	var server UDPServer
	raddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return nil, err
	}

	server.conn, err = net.ListenUDP("udp", raddr)
	if err != nil {
		return nil, err
	}

	return &server, nil
}

// ServeUDP starts the UDP server's read loop. The server should read from its
// listening socket and handle incoming client requests as according to the
// the specification.
func (u *UDPServer) handleCommand(b []byte, raddr *net.UDPAddr) error {
	result := "Unknown command"
	fields := strings.Split(string(b), "|:|")
	if len(fields) != 2 {
		goto send_response
	}
	switch c, s := fields[0], fields[1]; c {
	case "UPPER":
		result = strings.ToUpper(s)
	case "LOWER":
		result = strings.ToLower(s)
	case "CAMEL":
		result = strings.Title(strings.ToLower(s))
	case "ROT13":
		result = strings.Map(rot13, s)
	case "SWAP":
		result = strings.Map(swap, s)
	}

send_response:
	_, err := u.conn.WriteToUDP([]byte(result), raddr)
	return err
}

func rot13(r rune) rune {
	switch {
	case r >= 'A' && r <= 'Z':
		return 'A' + (r-'A'+13)%26

	case r >= 'a' && r <= 'z':
		return 'a' + (r-'a'+13)%26

	default:
		return r
	}
}
func swap(r rune) rune {
	switch {
	case unicode.IsLower(r):
		return unicode.ToUpper(r)
	case unicode.IsUpper(r):
		return unicode.ToLower(r)
	default:
		return r
	}
}

// socketIsClosed is a helper method to check if a listening socket has been
// closed.
func socketIsClosed(err error) bool {
	if strings.Contains(err.Error(), "use of closed network connection") {
		return true
	}
	return false
}
