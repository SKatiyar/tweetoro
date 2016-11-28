package tweetoro

import (
	"github.com/mrjones/oauth"
	"net/http"
)

type AuthConfig struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

func (ac AuthConfig) validate() error {
	if len(ac.ConsumerKey) < 1 {
		return ErrInvalidConsumerKey
	}
	if len(ac.ConsumerSecret) < 1 {
		return ErrInvalidConsumerSecret
	}
	if len(ac.AccessToken) < 1 {
		return ErrInvalidAccessToken
	}
	if len(ac.AccessTokenSecret) < 1 {
		return ErrInvalidAccessTokenSecret
	}

	return nil
}

func NewClient(params AuthConfig) (*http.Client, error) {
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
