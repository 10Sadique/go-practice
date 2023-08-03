package main

import "fmt"

func main() {
	fmt.Println(plus(1, 2))
	fmt.Println(plusPlus(1, 2, 3))

	num, str := vals()
	fmt.Println(num)
	fmt.Println(str)

	_, nstr := vals()
	fmt.Println(nstr)

	voidFunc("Jahan")
}

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func vals() (int, string) {
	return 3, "Another var"
}

func voidFunc(a string) {
	fmt.Printf("Hello, %s\n", a)
}
