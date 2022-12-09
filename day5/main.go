package main

import "fmt"

func greet(name string) {
	fmt.Println("Hello ", name)
}

func main() {
	go greet("Gopher")
	fmt.Println("Done")
}