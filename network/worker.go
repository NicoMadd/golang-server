package network

import (
	"fmt"
	"math/rand"
)

type Worker struct {
	id             int
	requestChannel chan Request
	routesMap      *RoutesMap
}

func (w Worker) String() string {
	return fmt.Sprintf("Worker { id: %d }", w.id)
}

func Init(channel chan Request, routesMap *RoutesMap) Worker {
	worker := Worker{requestChannel: channel, id: rand.Intn(10000), routesMap: routesMap}
	go worker.Handle()
	return worker
}

func (w Worker) Handle() {
	for {
		request := <-w.requestChannel

		request.OpenReader()

		request.ParseStartLine()

		route := w.routesMap.Get(request.method, request.path)

		response := NewResponse()

		if route == nil {
			response.NotFound()
		} else {
			route.Handler(request, response)
		}

		request.connection.Write(response.ToBytes())
		request.connection.Close()
	}
}
