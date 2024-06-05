package network

import (
	"fmt"
)

var (
	killerChannel chan bool = make(chan bool, 1)
)

const (
	WORKER_SIZE = 10
)

type Server struct {
	queue    chan Request
	workers  []Worker
	listener Listener
}

func (s *Server) String() string {
	// who all variables
	return fmt.Sprintf("Server { port: %d, queueSize: %d, workersSize: %v }", s.listener.port, cap(s.queue), len(s.workers))
}

func (s *Server) startListener() {
	go s.listener.StartListening()
}

func (s *Server) Start() error {
	fmt.Println("Server started")

	maxWorkers := cap(s.workers)

	// init workers
	for i := 0; i < maxWorkers; i++ {
		worker := Init(s.queue)
		s.workers = append(s.workers, worker)
	}

	// start listener thread
	s.startListener()

	<-killerChannel
	return nil
}

func InitHttpServer(port int16, workers_size int, request_buffer int) (*Server, error) {
	fmt.Printf("Initializing server on %d port...\n", port)

	// initialize requests channel
	requestsQueue := make(chan Request, request_buffer)

	// initialize workers
	workers := make([]Worker, workers_size)

	// initialize server
	server := Server{
		queue:   requestsQueue,
		workers: workers,
		// routes:  make(RoutesMap),
		listener: Listener{
			port:           port,
			requestChannel: requestsQueue,
		},
	}

	return &server, nil

}
