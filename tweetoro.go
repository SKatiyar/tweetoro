package tweetoro

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type FilterStreamOptions struct {
	AuthOpts      AuthConfig
	Delimited     string
	StallWarnings bool
	FilterLevel   string
	Language      []string
	Follow        []string
	Track         []string
	Locations     []float64
}

func (fso *FilterStreamOptions) ReqBody() (io.Reader, error) {
	data := url.Values{}
	data.Set("track", strings.Join(fso.Track, ","))

	return bytes.NewBufferString(data.Encode()), nil
}

type SampleStreamOptions struct {
	AuthOpts AuthConfig
}

func NewPublicFilterStream(opts FilterStreamOptions) (*Stream, error) {
	reqBody, reqBodyErr := opts.ReqBody()
	if reqBodyErr != nil {
		return nil, reqBodyErr
	}
	if paramsErr := opts.AuthOpts.validate(); paramsErr != nil {
		return nil, paramsErr
	}

	request, requestErr := http.NewRequest(http.MethodPost, PublicStreamFilterEndPoint, reqBody)
	if requestErr != nil {
		return nil, requestErr
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, responseErr := opts.AuthOpts.client().Do(request)
	if responseErr != nil {
		return nil, responseErr
	}

	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}

	return &Stream{response, bufio.NewScanner(response.Body)}, nil
}

func NewPublicSampleStream(opts SampleStreamOptions) (*Stream, error) {
	return &Stream{}, nil
}
