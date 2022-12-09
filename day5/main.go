package main

import "fmt"
import "time"

func greet(name string) {
	fmt.Println("Hello ", name)
}

func main() {
	go greet("Gopher")
	fmt.Println("Done")
	time.Sleep(time.Millisecond * 5)
}