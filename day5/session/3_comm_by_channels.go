package main

// Go Routine communication by Channel
// Sleep is not required, Locks are not required. Both should be handled by channel
// Blocked/Unbuffered channel works in a way that Producer needs to wait until consumer consumes the data. 
// So G2 waits until G1s work is done

import (
	"fmt"
	"math/rand"
)
func main() {
	ch := make(chan int)
	const n = 3
	for i:= 0; i<n; i++ {
		go func(){
			ch <- rand.Intn(100)
		}()
	}

	var data []int
	for i:= 0; i<n; i++ {
		data = append(data, <-ch)
	}

	fmt.Println(data)
}