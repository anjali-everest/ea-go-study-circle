package main

// Go ROutine communication by sharing memory
// Locking go routines

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)
func main() {
	var data []int

	const n = 3
	lock := sync.Mutex{}

	for i:= 0; i<n; i++ {
		go func(){
			lock.Lock()
			data = append(data, rand.Intn(50))
			lock.Unlock()
		}()
	}

	time.Sleep(30*time.Millisecond)
	fmt.Println(data)
}