package main

import (
	"fmt"
	"time"
)

func main() {
	doWork := func(strings []string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)
			for _, s := range strings {
				completed <- s
				fmt.Println(s)
			}
		}()
		return completed
	}

	jobs := []string{"job", "job2", "job3", "job4"}
	channel := doWork(jobs)

	for i := 0; i < len(jobs); i++ {
		fmt.Println(<-channel)
		time.Sleep(5 * time.Millisecond)
	}
	fmt.Println("Done")
}
