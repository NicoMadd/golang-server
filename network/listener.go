package network

import (
	"fmt"
	"golang-server/util"
	"net"
	"time"
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

	for {

		conn, err := listener.Accept()

		if err != nil {
			fmt.Println(err)
			return
		}

		go l.handleRequest(conn)
	}
}

func (l *Listener) status() string {
	return fmt.Sprintf("Listener { port: %d, reqCounter: %d }", l.port, l.reqCounter)
}

func (l *Listener) Monitor() string {
	for {
		time.Sleep(250 * time.Millisecond)
		util.ClearConsole()
		fmt.Println(l.status())

	}
}
