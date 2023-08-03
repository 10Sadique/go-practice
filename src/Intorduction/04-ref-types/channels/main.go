package main

import "fmt"

func main() {
	numCh := make(chan int, 10)

	for i := 1; i <= 10; i++ {
		numCh <- i
	}

	for j := 1; j <= 10; j++ {
		fmt.Println(<-numCh)
	}
}
