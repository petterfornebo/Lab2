// Leave an empty line above this comment.
package uecho

import (
	"net"
	"strings"
)

// UDPServer implements the UDP Echo Server specification found at
// https://github.com/COURSE_TAG/assignments/tree/master/lab2/README.md#udp-echo-server
type UDPServer struct {
	conn *net.UDPConn
	// TODO(student): Add fields if needed
}

// NewUDPServer returns a new UDPServer listening on addr. It should return an
// error if there was any problem resolving or listening on the provided addr.
func NewUDPServer(addr string) (*UDPServer, error) {
	// TODO(student): Implement
	return nil, nil
}

// ServeUDP starts the UDP server's read loop. The server should read from its
// listening socket and handle incoming client requests as according to the
// the specification.
func (u *UDPServer) ServeUDP() {
	// TODO(student): Implement
}

// socketIsClosed is a helper method to check if a listening socket has been
// closed.
func socketIsClosed(err error) bool {
	if strings.Contains(err.Error(), "use of closed network connection") {
		return true
	}
	return false
}
