package network

import (
	"bufio"
	"fmt"
	"golang-server/util"
	"net"
	"strconv"
	"strings"
)

type HTTPVersion string

type Path string

const (
	HTTP_0_9 HTTPVersion = "HTTP/0.9"
	HTTP_1_0 HTTPVersion = "HTTP/1.0"
	HTTP_1_1 HTTPVersion = "HTTP/1.1"
	HTTP_2_0 HTTPVersion = "HTTP/2.0"
	HTTP_3_0 HTTPVersion = "HTTP/3.0"
)

type HTTPMethod string

const (
	GET     HTTPMethod = "GET"
	POST    HTTPMethod = "POST"
	PUT     HTTPMethod = "PUT"
	DELETE  HTTPMethod = "DELETE"
	HEAD    HTTPMethod = "HEAD"
	OPTIONS HTTPMethod = "OPTIONS"
)

type Request struct {
	id          int64
	method      HTTPMethod
	path        Path
	httpVersion HTTPVersion
	headers     map[string]string
	size        int32
	body        []byte
	connection  net.Conn
	reader      *bufio.Reader
}

func InitRequest(conn net.Conn, counter int64) (Request, error) {
	return Request{
		id:          counter,
		httpVersion: "",
		headers:     make(map[string]string),
		size:        -1,
		body:        nil,
		connection:  conn,
		reader:      nil,
	}, nil
}

func (r *Request) OpenReader() error {
	r.reader = bufio.NewReader(r.connection)
	return nil
}

func (r *Request) StartLine() string {
	return fmt.Sprintf("%s %s %s", r.method, r.path, r.httpVersion)
}

func (r *Request) ParseStartLine() error {

	if r.httpVersion != "" {
		// already parsed
		return nil
	}

	readString, _ := r.reader.ReadString(util.LF)
	split_values := strings.Split(readString[:len(readString)-2], " ")

	r.method = HTTPMethod(split_values[0])
	r.path = Path(split_values[1])
	r.httpVersion = HTTPVersion(split_values[2])

	return nil
}

func (r *Request) Headers() string {
	return fmt.Sprintf("%s", r.headers)
}

func (r *Request) ParseHeaders() error {
	for {
		readString, _ := r.reader.ReadString(util.LF)
		if readString == "\r\n" {
			break
		}

		split_values := strings.Split(readString[:len(readString)-2], ": ")
		r.headers[split_values[0]] = split_values[1]
	}

	// get content length header value or set default value
	contentLength := r.headers["Content-Length"]
	if contentLength == "" {
		r.size = -1
	} else {
		value, _ := strconv.Atoi(contentLength)
		r.size = int32(value)
	}

	return nil
}

func (r *Request) Body() string {
	return fmt.Sprintf("%b", r.body)
}

func (r *Request) BodyAsString() string {
	// map byte array to string
	return string(r.body)
}

func (r *Request) ParseBody() error {

	// read content length size from body
	if r.size == -1 {
		return nil
	}

	r.body = make([]byte, r.size)
	_, err := r.reader.Read(r.body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (r *Request) ParseRequest() error {

	_ = r.ParseStartLine()

	_ = r.ParseHeaders()

	_ = r.ParseBody()

	return nil

}
