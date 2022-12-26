package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	newRandStream := func(streamLength int) <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)
			for i := 0; i < streamLength; i++ {
				randStream <- rand.Int()
			}
		}()
		return randStream
	}

	streamLength := 3
	randStream := newRandStream(streamLength)
	fmt.Printf("%d random ints:\n", streamLength)
	for i := 1; i <= streamLength; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
		time.Sleep(5 * time.Millisecond)
	}

}
