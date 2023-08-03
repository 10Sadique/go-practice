package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	sum := 0
	for i, num := range nums {
		if num == 3 {
			fmt.Println("Index: ", i)
		}
	}

	fmt.Println(sum)
}