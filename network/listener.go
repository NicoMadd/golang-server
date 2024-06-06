package network

import (
	"fmt"
	"net"
)

const (
	Localhost = "127.0.0.1"
	TCP       = "tcp"
)

type Listener struct {
	port           int16
	requestChannel chan<- Request
	reqCounter     int64
}

func (l *Listener) handleRequest(conn net.Conn) {
	request, err := InitRequest(conn, l.reqCounter)
	l.reqCounter++

	if l.reqCounter%1000 == 0 {
		fmt.Println("Request counter: ", l.reqCounter)
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	l.requestChannel <- request

}

func (l *Listener) StartListening() {

	// open socket
	listener, err := net.Listen(TCP, fmt.Sprintf("%s:%d", Localhost, l.port))

	fmt.Printf("Listening on port %d...\n", l.port)

	if err != nil {
		fmt.Println(err)
		return
	}

	reqCounter := 0

	for {
		fmt.Println("listening...", reqCounter)
		reqCounter++

		conn, err := listener.Accept()

		if err != nil {
			fmt.Println(err)
			return
		}

		go l.handleRequest(conn)
	}
}
