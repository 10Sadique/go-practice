package main

import (
	"fmt"
	"reflect"
)

func main() {
	rune1 := 'B'
	rune2 := 'g'
	rune3 := '\a'

	fmt.Printf("\nRene 1: %c; Unicode: %U; Type: %s", rune1, rune1, reflect.TypeOf(rune1))
	fmt.Printf("\nRene 2: %c; Unicode: %U; Type: %s", rune2, rune2, reflect.TypeOf(rune2))
	fmt.Printf("\nRene 3: %c; Unicode: %U; Type: %s", rune3, rune3, reflect.TypeOf(rune3))
}