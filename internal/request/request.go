package request

import (
	"bytes"
	"fmt"
	"io"
	"demoproject/internal/common"
	
)

type ParserState string

const (
	stateInit ParserState = "init"
	stateDone ParserState = "done"
	stateError ParserState = "error"
)

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

type Request struct {
	RequestLine RequestLine
	state       ParserState
}

func newRequest() *Request {
	return &Request{
		state: stateInit,
	}

}

var ErrorMalformedRequestLine = fmt.Errorf("malformed request-line")
var ErrorUnsupportedHttpVersion = fmt.Errorf("unsupported http version")
var ErrorRequestInErrorState = fmt.Errorf("request in Error state")

var SEPERATOR = []byte("\r\n") // got abysmal amount of errors becasue i put /r/n instead of \r\n Was just about to rage quit

func parseRequestLine(b []byte) (*RequestLine, int, error) {


	idx := bytes.Index(b, SEPERATOR)
	if idx == -1 {
		return nil, 0, nil
	}

	startLine := b[:idx]
	read := idx+len(SEPERATOR)

	parts := bytes.Split(startLine, []byte(" "))
	if len(parts) != 3 {
		return nil, 0, ErrorMalformedRequestLine
	}

	httpsParts := bytes.Split(parts[2], []byte("/"))
	if len(httpsParts) != 2 || string(httpsParts[0]) != "HTTP" || string(httpsParts[1]) != "1.1" {
		return nil, 0, ErrorMalformedRequestLine
	}

	rl := &RequestLine{
		Method:        string(parts[0]),
		RequestTarget: string(parts[1]),
		HttpVersion:   string(httpsParts[1]),
	}

	return rl, read, nil
}

func (r *Request) parse(data []byte) (int, error) {

	common.WriteLog(string(data), "./../../Log.txt")


	read := 0
	outer: 
	for {
		switch r.state {
		case stateError:
			return 0, ErrorRequestInErrorState

		case stateInit:
			rl, n, err := parseRequestLine(data[read:])
			if err != nil {
				r.state = stateError
				return 0, err
			}

			if n == 0 {
				break outer
			}

			r.RequestLine = *rl
			read += n

			r.state = stateDone

		case stateDone:
			break outer
		}
	}
	return read, nil
}

func (r *Request) done() bool {
	return r.state == stateDone || r.state == stateError
}



func RequestFromReader(reader io.Reader) (*Request, error) {
	request := newRequest()

	buf := make([]byte, 1024)
	bufLen := 0
	for !request.done() {
		n, err := reader.Read(buf[bufLen:])
		// TODO: What to do here??
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
