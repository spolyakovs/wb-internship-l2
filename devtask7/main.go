package main

import (
	"fmt"
	"time"
)

func main() {
	var or func(channels ...<-chan interface{}) <-chan interface{} = orChannels

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("Done after %v\n", time.Since(start))

}

func orChannels(channels ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})

	for i := range channels {
		go func(channel <-chan interface{}) {
			for range channel {
			}
			out <- true
		}(channels[i])
	}

	return out
}
