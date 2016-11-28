package main

import (
	"github.com/SKatiyar/tweetoro"
	"log"
)

func main() {
	stream, streamErr := tweetoro.NewPublicFilterStream(tweetoro.FilterStreamOptions{})
	if streamErr != nil {
		log.Println(streamErr.Error())
		return
	}

	var data interface{}
	for stream.Next(&data) {
		log.Println(data)
	}
}
