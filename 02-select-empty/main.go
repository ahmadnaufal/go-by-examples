package main

import "fmt"

func PrintExample() {
	fmt.Println("Hello from goroutines")
}

// example usage of select in goroutines
func main() {
	fmt.Println("Hello from main thread")

	go PrintExample()

	// this line below will result in:
	// fatal error: all goroutines are asleep - deadlock!

	// goroutine 1 [select (no cases)]:
	// main.main()
	// 				/Users/fahmadnaufal/projects/go-by-examples/02-select-empty/main.go:15 +0x68
	// exit status 2
	select {}

	fmt.Println("Complete!")
}
