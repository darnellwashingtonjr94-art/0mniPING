package mocks

import (
	"net"
)

func NewTestTCPServer() (string, func()) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			go func(c net.Conn) {
				defer c.Close()
				buf := make([]byte, 512)
				c.Read(buf)
				c.Write([]byte("ACK"))
			}(conn)
		}
	}()

	return listener.Addr().String(), func() {
		listener.Close()
	}
}
