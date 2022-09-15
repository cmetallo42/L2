package main

import (
	"fmt"
	"sync"
	"time"
)

func merge(channels ...<-chan any) <-chan any {
	out := make(chan any)
	go func() {
		var wg sync.WaitGroup
		wg.Add(len(channels))

		for i, ch := range channels {
			go func(ch <-chan any, i int) {
				for v := range ch {
					out <- v
				}
				fmt.Printf("Channel %v is Done\n", i)
				wg.Done()
			}(ch, i)
		}
		wg.Wait()
		fmt.Println("All channels is closed!")
		close(out)
	}()
	return out
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-merge(
		sig(1*time.Second),
		sig(2*time.Second),
		sig(3*time.Second),
		sig(4*time.Second),
		sig(5*time.Second),
	)

	fmt.Printf("Done after %v\n", time.Since(start))
}
