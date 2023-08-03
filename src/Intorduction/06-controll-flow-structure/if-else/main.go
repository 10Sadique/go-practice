package main

import "fmt"

func main() {
	isEven(7)
	isEven(8)

	isDivisibleBy4(8)
}

func isEven(n int) {
	if n%2 == 0 {
		fmt.Printf("%d is even.\n", n)
	} else {
		fmt.Printf("%d is odd.\n", n)
	}
}

func isDivisibleBy4(n int) {
	if n%4==0 {
		fmt.Printf("%d is divisible by 4\n", n)
	}
}