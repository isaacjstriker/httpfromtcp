package request

import (
	"io"
	"strings"
	"errors"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method		  string
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	reqStr := string(data)

	crlfIdx := strings.Index(reqStr, "\r\n")
	if crlfIdx == -1 {
		return nil, errors.New("malformed request: missing CRLF after request-line")
	}

	line := reqStr[:crlfIdx]

	rl, err := parseRequesetLine(line)
	if err != nil {
		return nil, err
	}

	return &Request{RequestLine: *rl}, nil
}

func parseRequesetLine(line string) (*RequestLine, error) {
    parts := strings.Split(line, " ")
    if len(parts) != 3 {
        return nil, errors.New("malformed request-line: expected 3 parts")
    }

    method := parts[0]
    target := parts[1]
    versionToken := parts[2]

    // Validate method: only capital A-Z letters.
    if method == "" {
        return nil, errors.New("empty method")
    }
    for _, r := range method {
        if r < 'A' || r > 'Z' {
            return nil, errors.New("invalid method: must be A-Z only")
        }
    }

    // Validate HTTP version: must be HTTP/1.1 and we store just "1.1".
    if !strings.HasPrefix(versionToken, "HTTP/") {
        return nil, errors.New("invalid HTTP version prefix")
    }
    version := strings.TrimPrefix(versionToken, "HTTP/")
    if version != "1.1" {
        return nil, errors.New("unsupported HTTP version (only 1.1 supported)")
    }

    return &RequestLine{
        HttpVersion:   version,
        RequestTarget: target,
        Method:        method,
    }, nil
}