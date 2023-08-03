package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:4]

	fmt.Println("array: ", array)
	fmt.Println("slice: ", slice)
	fmt.Println(len(slice))
}
