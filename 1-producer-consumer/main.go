//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer scenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)

func producer(tweets chan *Tweet, stream Stream) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			close(tweets)
			return
		}

		tweets <- tweet
	}
}

func consumer(tweets chan *Tweet, done chan<- struct{}) {
	for {
		tweet, ok := <-tweets
		if !ok {
			done <- struct{}{}
			return
		}

		if tweet.IsTalkingAboutGo() {
			fmt.Println(tweet.Username, "\ttweets about golang")
		} else {
			fmt.Println(tweet.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()
	tweets := make(chan *Tweet)
	done := make(chan struct{})

	// Producer
	go producer(tweets, stream)

	// Consumer
	go consumer(tweets, done)

	<-done

	fmt.Printf("Process took %s\n", time.Since(start))
}
