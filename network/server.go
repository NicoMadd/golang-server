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
	routes   *RoutesMap
}

func (s *Server) String() string {
	// who all variables
	return fmt.Sprintf("Server { port: %d, queueSize: %d, workersSize: %v }", s.listener.port, cap(s.queue), len(s.workers))
}

func (s *Server) Start() error {
	fmt.Println("Server started")

	maxWorkers := cap(s.workers)

	// init workers
	for i := 0; i < maxWorkers; i++ {
		worker := Init(s.queue, s.routes)
		s.workers = append(s.workers, worker)
	}

	// start listener thread
	go s.listener.StartListening()

	// start listener monitor
	go s.listener.Monitor()

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
		routes:  new(RoutesMap),
		listener: Listener{
			port:           port,
			requestChannel: requestsQueue,
		},
	}

	return &server, nil

}

func (s *Server) Get(path Path, handler Handler) {
	s.routes.AddHandler("GET", path, handler)
}

func (s *Server) Post(path Path, handler Handler) {
	s.routes.AddHandler("POST", path, handler)
}

func (s *Server) Put(path Path, handler Handler) {
	s.routes.AddHandler("PUT", path, handler)
}

func (s *Server) Delete(path Path, handler Handler) {
	s.routes.AddHandler("DELETE", path, handler)
}

func (s *Server) Head(path Path, handler Handler) {
	s.routes.AddHandler("HEAD", path, handler)
}

func (s *Server) Options(path Path, handler Handler) {
	s.routes.AddHandler("OPTIONS", path, handler)
}
