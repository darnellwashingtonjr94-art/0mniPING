package mocks

import (
	"net"
)

func NewTestUDPServer() (string, func()) {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 0})
	if err != nil {
		panic(err)
	}

	go func() {
		buf := make([]byte, 1024)
		for {
			n, remote, err := conn.ReadFromUDP(buf)
			if err != nil {
				return
			}
			conn.WriteToUDP(buf[:n], remote)
		}
	}()

	return conn.LocalAddr().String(), func() {
		conn.Close()
	}
}
