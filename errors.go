package tweetoro

import (
	"errors"
)

var (
	ErrInvalidConsumerKey       = errors.New("Invalid consumer key")
	ErrInvalidConsumerSecret    = errors.New("Invalid consumer secret")
	ErrInvalidAccessToken       = errors.New("Invalid access token")
	ErrInvalidAccessTokenSecret = errors.New("Invalid access token secret")
)
