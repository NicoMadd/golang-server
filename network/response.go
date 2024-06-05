package network

import "fmt"

type Response struct {
	httpVersion  HTTPVersion
	statusCode   int
	reasonPhrase string
}

func (response Response) String() string {
	return fmt.Sprintf("%s %d %s", response.httpVersion, response.statusCode, response.reasonPhrase)
}
