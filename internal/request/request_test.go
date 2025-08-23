package request

import (
	"testing"
	"strings"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRequestLineParse(t *testing.T) {
	r, err := RequestFromReader(strings.NewReader("GET / HTTP/1.1\r\nHost: localhost:42069\r\nUser-Agent: curl/7.81.0\r\nAccept: */*\r\n\r\n"))
	require.NoError(t, err)
	require.NotNil(t, r)
	assert.Equal(t, "GET", r.RequestLine.Method)
	assert.Equal(t, "/", r.RequestLine.RequestTarget)
	assert.Equal(t, "1.1", r.RequestLine.HttpVersion)

	// Test: Good GET Request line with path
	r, err = RequestFromReader(strings.NewReader("GET /coffee HTTP/1.1\r\nHost: localhost:42069\r\nUser-Agent: curl/7.81.0\r\nAccept: */*\r\n\r\n"))
	require.NoError(t, err)
	require.NotNil(t, r)
	assert.Equal(t, "GET", r.RequestLine.Method)
	assert.Equal(t, "/coffee", r.RequestLine.RequestTarget)
	assert.Equal(t, "1.1", r.RequestLine.HttpVersion)

	// Test: Invalid number of parts in request line
	_, err = RequestFromReader(strings.NewReader("/coffee HTTP/1.1\r\nHost: localhost:42069\r\nUser-Agent: curl/7.81.0\r\nAccept: */*\r\n\r\n"))
	require.Error(t, err)

    // Test: Good POST Request with path
    r, err = RequestFromReader(strings.NewReader("POST /submit HTTP/1.1\r\nHost: localhost\r\n\r\n"))
    require.NoError(t, err)
    assert.Equal(t, "POST", r.RequestLine.Method)
    assert.Equal(t, "/submit", r.RequestLine.RequestTarget)

    // Test: Invalid method (out of order) Request line (lowercase letter)
    _, err = RequestFromReader(strings.NewReader("GeT / HTTP/1.1\r\nHost: localhost\r\n\r\n"))
    require.Error(t, err)

    // Test: Invalid method (non A-Z character)
    _, err = RequestFromReader(strings.NewReader("GET1 / HTTP/1.1\r\nHost: localhost\r\n\r\n"))
    require.Error(t, err)

    // Test: Invalid version in Request line (HTTP/2.0)
    _, err = RequestFromReader(strings.NewReader("GET / HTTP/2.0\r\nHost: localhost\r\n\r\n"))
    require.Error(t, err)

    // Test: Invalid version in Request line (HTTP/1.0)
    _, err = RequestFromReader(strings.NewReader("GET / HTTP/1.0\r\nHost: localhost\r\n\r\n"))
    require.Error(t, err)

    // Test: Invalid version token missing HTTP/ prefix
    _, err = RequestFromReader(strings.NewReader("GET / 1.1\r\nHost: localhost\r\n\r\n"))
    require.Error(t, err)

    // Test: Missing CRLF after request-line
    _, err = RequestFromReader(strings.NewReader("GET / HTTP/1.1")) // no \r\n
    require.Error(t, err)

    // Test: Multiple spaces producing too many parts
    _, err = RequestFromReader(strings.NewReader("GET  / HTTP/1.1\r\nHost: localhost\r\n\r\n"))
    require.Error(t, err)

    // Test: Empty method
    _, err = RequestFromReader(strings.NewReader(" / HTTP/1.1\r\nHost: localhost\r\n\r\n"))
    require.Error(t, err)

    // Test: Path with query string (should succeed)
    r, err = RequestFromReader(strings.NewReader("GET /coffee?milk=yes HTTP/1.1\r\nHost: localhost\r\n\r\n"))
    require.NoError(t, err)
    assert.Equal(t, "/coffee?milk=yes", r.RequestLine.RequestTarget)
}