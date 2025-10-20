package stringRequest

import (
	"demoproject/internal/common"
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

var ErrorMalformedRequestLine = fmt.Errorf("malformed request-line")
var ErrorUnsupportedHttpVersion = fmt.Errorf("unsupported http version")
var SEPERATOR = "\r\n"


func parseRequestLine(b string) (*RequestLine, string, error){

	common.WriteLog(string(b), "./../../Test.log")


	idx := strings.Index(b, SEPERATOR)
	if idx == -1 {
		return nil, b, nil
	}

	startLine := b[:idx]
	restOfMsg := b[idx+len(SEPERATOR):]

	parts := strings.Split(startLine, " ")
	if len(parts) != 3 {
		return nil, restOfMsg, ErrorMalformedRequestLine
	}

	httpsParts := strings.Split(parts[2], "/")
	if len(httpsParts) != 2 || httpsParts[0] != "HTTP" || httpsParts[1] != "1.1" {
		return nil, restOfMsg, ErrorMalformedRequestLine
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

	common.WriteLog(string(data), "./../../Test.log")


	str := string(data)

	rl, _, err := parseRequestLine(str)
	if err != nil {
		return nil, err
	}

	return &Request{
		RequestLine: *rl,
	}, err

}