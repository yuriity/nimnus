package request

import (
	"bytes"
	"fmt"
	"io"
)

type parserState string
const (
	Initialized parserState = "initialized"
	Done parserState = "done"
)

type RequestLine struct {
	Method        string
	RequestTarget string
	HttpVersion   string
}

type Request struct {
	RequestLine RequestLine
	state parserState
}

func newRequest() *Request {
	return &Request{
		state: Initialized,
	}
}

var ERROR_MALFORMED_REQUEST_LINE = fmt.Errorf("malformed request-line")
var ERROR_UNSOPPORTED_HTTP_VERSION = fmt.Errorf("unsupported HTTP version")
var ERROR_INVALID_METHOD = fmt.Errorf("invalid method")
var ERROR_EMPTY_METHOD = fmt.Errorf("empty method")
var ERROR_UNKNOWN_STATE = fmt.Errorf("unknown state")
var ERROR_ALREADY_DONE = fmt.Errorf("trying to read data in a done state")
var SEPARATOR = []byte("\r\n")

func parseRequestLine(b []byte) (*RequestLine, int, error) {
	idx := bytes.Index(b, SEPARATOR)
	if idx == -1 {
		return nil, 0, nil
	}

	startLine := b[:idx]
	read := idx + len(SEPARATOR)

	parts := bytes.Split(startLine, []byte(" "))
	if len(parts) != 3 {
		return nil, 0, ERROR_MALFORMED_REQUEST_LINE
	}

	httpParts := bytes.Split(parts[2], []byte("/"))
	if len(httpParts) != 2 || string(httpParts[0]) != "HTTP" || string(httpParts[1]) != "1.1" {
		return nil, 0, ERROR_MALFORMED_REQUEST_LINE
	}

	rl := &RequestLine{
		Method: string(parts[0]),
		RequestTarget: string(parts[1]),
		HttpVersion: string(httpParts[1]),
	}

	return rl, read, nil
}

func (r *Request) parse(data []byte) (int, error) {
	read := 0

outer:
	for {
		switch r.state {
		case Initialized:
			rl, n, err := parseRequestLine(data[read:])
			if err != nil {
				return 0, err
			}
			if n == 0 {
				break outer
			}

			r.RequestLine = *rl
			read += n

			r.state = Done

		case Done:
			break outer
		default:
			return 0, ERROR_UNKNOWN_STATE
		}
	}

	return read, nil
}

func (r *Request) done() bool {
	return r.state == Done
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	request := newRequest()

	// NOTE: buffer could get overrun... a header that exceeds 1k would do that...
	// or body
	buf := make([]byte, 1024)
	bufLen := 0
	for !request.done() {
		n, err := reader.Read(buf[bufLen:])
		// TODO: what to do here?
		if err != nil {
			return nil, err
		}

		bufLen += n
		readN, err := request.parse(buf[:bufLen])
		if err != nil {
			return nil, err
		}

		copy(buf, buf[readN:bufLen])
		bufLen -= readN
	}

	return request, nil
}
