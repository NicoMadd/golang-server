package main

import (
	"fmt"
	"server/network"
)

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

	err = server.Start()

	if err != nil {
		fmt.Println(err)
		return
	}

}
