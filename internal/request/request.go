package request

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

type Request struct {
	RequestLine RequestLine
}

var ERROR_MALFORMED_REQUEST_LINE = fmt.Errorf("malformed request-line")
var ERROR_UNSUPPORTED_HTTP_VERSION = fmt.Errorf("unsupported http version")
var SEPERATOR = "\r\n" // got abysmal amount of errors becasue i put /r/n instead of \r\n Was just about to rage quit


func parseRequestLine(b string) (*RequestLine, string, error){

	idx := strings.Index(b, SEPERATOR)
	if idx == -1 {
		return nil, b, nil
	}

	startLine := b[:idx]
	restOfMsg := b[idx+len(SEPERATOR):]

	parts := strings.Split(startLine, " ")
	if len(parts) != 3 {
		return nil, restOfMsg, ERROR_MALFORMED_REQUEST_LINE
	}

	httpsParts := strings.Split(parts[2], "/")
	if len(httpsParts) != 2 || httpsParts[0] != "HTTP" || httpsParts[1] != "1.1" {
		return nil, restOfMsg, ERROR_MALFORMED_REQUEST_LINE
	}

	rl := &RequestLine{
		Method: parts[0],
		RequestTarget: parts[1],
		HttpVersion: httpsParts[1],
	} 

	return rl, restOfMsg, nil
}

func RequestFromReader(reader io.Reader) (*Request, error) {

	data, err := io.ReadAll(reader) 
	if err != nil {
		return nil, errors.Join(
			fmt.Errorf("unable to io.ReadAll"), err,
		)
	}

	str := string(data)

	rl, _, err := parseRequestLine(str)
	if err != nil {
		return nil, err
	}

	return &Request{
		RequestLine: *rl,
	}, err

}