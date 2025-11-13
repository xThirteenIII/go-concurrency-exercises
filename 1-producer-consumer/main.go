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

func producer(stream Stream, tweetCh chan<- *Tweet) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			close(tweetCh)
			return
		}
		tweetCh <- tweet
	}
}

func consumer(ch <-chan *Tweet) {
	for t := range ch {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()
	tweetsCh := make(chan *Tweet)
	var wg sync.WaitGroup

	// Producer
	wg.Go(func() { producer(stream, tweetsCh) })

	// Consumer
	wg.Go(func() { consumer(tweetsCh) })

	wg.Wait()
	fmt.Printf("Process took %s\n", time.Since(start))
}
