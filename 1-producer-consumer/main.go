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
	"sync"
	"time"
)

func producer(tweets chan *Tweet, wg *sync.WaitGroup, stream Stream) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			wg.Done()
			return
		}

		wg.Add(1)
		tweets <- tweet
	}
}

func consumer(tweets chan *Tweet, wg *sync.WaitGroup) {
	for tweet := range tweets {
		if tweet.IsTalkingAboutGo() {
			fmt.Println(tweet.Username, "\ttweets about golang")
		} else {
			fmt.Println(tweet.Username, "\tdoes not tweet about golang")
		}
		wg.Done()
	}
}

func main() {
	start := time.Now()
	tweets := make(chan *Tweet)
	stream := GetMockStream()
	wg := &sync.WaitGroup{}

	wg.Add(1)
	// Producer
	go producer(tweets, wg, stream)

	// Consumer
	go consumer(tweets, wg)
	wg.Wait()

	fmt.Printf("Process took %s\n", time.Since(start))
}
