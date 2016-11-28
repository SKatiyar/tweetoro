package tweetoro

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

type SampleStreamOptions struct {
	AuthOpts AuthConfig
}

func NewPublicFilterStream(opts FilterStreamOptions) (*Stream, error) {
	client, clientErr := NewClient(opts.AuthOpts)
	if clientErr != nil {
		return nil, clientErr
	}

	response, responseErr := client.Post(PublicStreamFilterEndPoint, "application/json", nil)
	if responseErr != nil {
		return nil, responseErr
	}

	return &Stream{response}, nil
}

func NewPublicSampleStream(opts SampleStreamOptions) (*Stream, error) {
	return &Stream{}, nil
}
