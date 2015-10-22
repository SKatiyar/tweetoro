package tweetoro

import (
	"github.com/mrjones/oauth"
	"net/http"
)

const (
	StreamEndPoint string = "https://stream.twitter.com/1.1/statuses/filter.json"
)

type Config struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

func (c Config) validate() error {
	if len(c.ConsumerKey) < 1 {
		return ErrInvalidConsumerKey
	}
	if len(c.ConsumerSecret) < 1 {
		return ErrInvalidConsumerSecret
	}
	if len(c.AccessToken) < 1 {
		return ErrInvalidAccessToken
	}
	if len(c.AccessTokenSecret) < 1 {
		return ErrInvalidAccessTokenSecret
	}

	return nil
}

func NewClient(params Config) (*http.Client, error) {
	if paramsErr := params.validate(); paramsErr != nil {
		return nil, paramsErr
	}

	auth := oauth.NewConsumer(params.ConsumerKey, params.ConsumerSecret, oauth.ServiceProvider{})

	client, clientErr := auth.MakeHttpClient(&oauth.AccessToken{
		Token:  params.AccessToken,
		Secret: params.AccessTokenSecret,
	})
	if clientErr != nil {
		return nil, clientErr
	}

	return client, nil
}
