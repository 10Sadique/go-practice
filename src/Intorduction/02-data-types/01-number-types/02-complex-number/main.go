package main

import "fmt"

func main() {
	c1 := complex(2, 4)
	c2 := 4 + 5i
	c3 := c1 + c2

	fmt.Println("ADD: ", c3)
	
	re:= real(c3)
	im:= imag(c3)

	fmt.Println(re, im)
}