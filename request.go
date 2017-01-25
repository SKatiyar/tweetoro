package tweetoro

import (
	"errors"
	"io"
	"net/http"
)

const (
	PublicStreamFilterEndPoint string = "https://stream.twitter.com/1.1/statuses/filter.json"
	PublicStreamSampleEndPoint string = "https://stream.twitter.com/1.1/statuses/sample.json"
)

func NewClient(conf AuthConfig, url string, method string, body io.Reader) (*http.Response, error) {
	if paramsErr := conf.validate(); paramsErr != nil {
		return nil, paramsErr
	}

	request, requestErr := http.NewRequest(method, url, body)
	if requestErr != nil {
		return nil, requestErr
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, responseErr := conf.client().Do(request)
	if responseErr != nil {
		return nil, responseErr
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}

	return response, nil
}
