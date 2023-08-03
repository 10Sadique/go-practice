package main

import "fmt"

func main()  {
	a:=10
	b:=&a

	fmt.Println(b)
	fmt.Println(*b)

	*b= 200

	fmt.Println(b)
	fmt.Println(*b)
}