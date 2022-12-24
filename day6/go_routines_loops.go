package main

import "fmt"

func execute(index int, finished chan bool) {
	fmt.Println("I am the", index, "th goroutine")
	finished <- true
}

func main() {

	done := make(chan bool)

	// create and spawn ten goroutines
	for i := 0; i < 10; i++ {
		// define a closure and immediately spawn it as a goroutine.
		// Pass the loop index (new!) and the done channel
		go execute(i, done) // now we also pass i to the goroutine
	}

	// wait for the goroutines to finish
	for i := 0; i < 10; i++ {
		<-done
	}

	fmt.Println("Done.")
}
