package tweetoro

import (
	"net/http"
)

type Stream struct {
	resStream *http.Response
}

func (s *Stream) Next(dst interface{}) bool {
	return false
}

func (s *Stream) Error() error {
	return nil
}

func (s *Stream) Close() error {
	return nil
}
