package main

import "fmt"

func PrintExample(c chan int) {
	fmt.Println("Hello from goroutines")
	c <- 100
}

// example usage of select in goroutines
func main() {
	c := make(chan int)
	fmt.Println("Hello from main thread")

	go PrintExample(c)

	// select {
	// case <-c:
	// 	fmt.Println("Hello!", c)
	// }

	fmt.Println(<-c)

	fmt.Println("Complete!")
}
