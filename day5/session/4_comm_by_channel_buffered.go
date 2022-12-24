package main

// Go Routine communication by Channel
// Sleep is not required, Locks are not required. Both should be handled by channel
// buffered channel works in a way that Producer can produce as many but it has to wait until it becomes empty.
// consumer doesn't need to wait it can consumes the data. 
// So G2 waits until G1s work is done

//Channel size 0 is default, works as blocked/non-buffered channel

import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
	ch := make(chan int, 4)
	const n = 10
	for i:= 0; i<n; i++ {
		go func(){
			time.Sleep(10*time.Millisecond)
			ch <- rand.Intn(100)
			fmt.Println("Pushed")
		}()
	}

	var data []int
	for i:= 0; i<n; i++ {
		time.Sleep(30*time.Millisecond)
		data = append(data, <-ch)
		fmt.Println("Read")
	}

	fmt.Println(data)
}