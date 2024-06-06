package network

import "fmt"

type Response struct {
	httpVersion  HTTPVersion
	statusCode   int
	reasonPhrase string
}

func NewResponse() *Response {
	return &Response{
		httpVersion:  HTTP_1_1,
		statusCode:   200,
		reasonPhrase: "OK",
	}
}

func (response Response) String() string {
	return fmt.Sprintf("%s %d %s\r\n\r\n", response.httpVersion, response.statusCode, response.reasonPhrase)
}

func (response Response) ToBytes() []byte {
	return []byte(response.String())
}

func (response *Response) SetStatusCode(statusCode int) {
	response.statusCode = statusCode
}
func (response *Response) OK() {
	response.statusCode = 200
}

func (response *Response) NotFound() {
	response.statusCode = 404
}

func (response *Response) SetReasonPhrase(reasonPhrase string) {
	response.reasonPhrase = reasonPhrase
}
