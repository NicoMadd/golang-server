package main

import (
	"fmt"
	"golang-server/network"
)

func setupRoutes(server *network.Server) {
	fmt.Println("Setting up routes...")

	server.Get("/hello", func(request network.Request, response *network.Response) {

		response.SetReasonPhrase("asd")

	})

	server.Post("/hello", func(request network.Request, response *network.Response) {
		fmt.Println("Hello, world")
	})
}

func main() {

	fmt.Println("Initializing...")

	var port int16 = 8081
	var workers_size int = 10
	var err error = nil

	server, err := network.InitHttpServer(port, workers_size, 10)

	fmt.Println(server.String())

	if err != nil {
		fmt.Println(err)
		return
	}

	setupRoutes(server)

	err = server.Start()

	if err != nil {
		fmt.Println(err)
		return
	}

}
