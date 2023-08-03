package main

import (
	"fmt"
)

func main()  {
	var x float64 = 20.0
	fmt.Println(x);
	fmt.Printf("x is a type of %T\n", x);

	var num int
	var amount float32
	var isValid bool
	var name string

	num = 10;
	amount = 49.99
	isValid = true
	name = "Jahan"

	fmt.Println(num, amount, isValid, name)

	// declaring w/o data type 
	var num2 = 20
	var amount2 = 49.99
	var isValid2 = true
	var name2 = "Jahan"

	fmt.Println(num2, amount2, isValid2, name2)

	// shorthand
	y := 42
	fmt.Println(y)
	a, b := 1, 2
	fmt.Println(a, b)

	var1, var2 :=1, "hi"
	fmt.Println(var1, var2)

	var3, var2 := 2, "hello"
	fmt.Println(var3, var2)

	
}