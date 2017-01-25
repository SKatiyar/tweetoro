package tweetoro

import (
	"bufio"
	"bytes"
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

	response, responseErr := NewClient(opts.AuthOpts, PublicStreamFilterEndPoint, http.MethodPost, reqBody)
	if responseErr != nil {
		return nil, responseErr
	}

	newScanner := bufio.NewScanner(response.Body)
	newScanner.Split(bufio.SplitFunc(ScanCRLF))

	return &Stream{response, newScanner}, nil
}

func NewPublicSampleStream(opts SampleStreamOptions) (*Stream, error) {
	return &Stream{}, nil
}
