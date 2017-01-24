package main

import (
	"github.com/SKatiyar/tweetoro"
	"log"
)

func main() {
	stream, streamErr := tweetoro.NewPublicFilterStream(tweetoro.FilterStreamOptions{})
	if streamErr != nil {
		log.Println("1", streamErr.Error())
		return
	}

	for stream.Next() {
		if streamErr := stream.Error(); streamErr != nil {
			log.Println("2", streamErr)
			break
		}

		var data interface{}
		if decodeErr := stream.Scan(&data); decodeErr != nil {
			log.Println("3", decodeErr)
			break
		}

		log.Println("4", data)
	}

	log.Println("5", stream.Close())
}
