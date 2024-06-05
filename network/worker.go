package network

import (
	"fmt"
	"math/rand"
)

type Worker struct {
	id             int
	requestChannel chan Request
}

func (w Worker) String() string {
	return fmt.Sprintf("Worker { id: %d }", w.id)
}

func Init(channel chan Request) Worker {
	worker := Worker{requestChannel: channel, id: rand.Intn(10000)}
	go worker.Handle()
	return worker
}

func (w Worker) Handle() {

	for {
		request := <-w.requestChannel

		request.ParseRequest()

		request.connection.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
		request.connection.Close()
	}
}
