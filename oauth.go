package tweetoro

import (
	"github.com/dghubble/oauth1"
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

func (ac AuthConfig) client() *http.Client {
	return oauth1.NewConfig(ac.ConsumerKey, ac.ConsumerSecret).
		Client(oauth1.NoContext, oauth1.NewToken(ac.AccessToken, ac.AccessTokenSecret))
}
