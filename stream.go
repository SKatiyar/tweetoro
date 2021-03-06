package tweetoro

import (
	"bufio"
	"bytes"
	"encoding/json"
	"net/http"
)

type Stream struct {
	resStream *http.Response
	scanner   *bufio.Scanner
}

func (s *Stream) Next() bool {
	return s.scanner.Scan()
}

func (s *Stream) Scan(dst interface{}) error {
	data := s.scanner.Bytes()
	if len(data) != 0 {
		return json.Unmarshal(data, dst)
	}

	return nil
}

func (s *Stream) Error() error {
	return s.scanner.Err()
}

func (s *Stream) Close() error {
	return s.resStream.Body.Close()
}

func ScanCRLF(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, make([]byte, 0, 1), nil
	}
	if i := bytes.Index(data, []byte{'\r', '\n'}); i >= 0 {
		// We have a full newline-terminated line.
		println(i)
		return i + 2, dropCR(data[0:i]), nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), dropCR(data), nil
	}
	// Request more data.
	return 0, make([]byte, 0, 1), nil
}

// dropCR drops a terminal \r from the data.
func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}
