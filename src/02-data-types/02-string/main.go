package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	var str string = "Hello, World"
	fmt.Println(str)

	str1:= "Hello, "
	str2:="World"

	fmt.Println(str1+str2)

	str = "Hello, World!"
	fmt.Println(str[0:5])

	fmt.Println(str[8])
	fmt.Println(len(str))
	fmt.Println(strings.ToLower(str))
	fmt.Println(strings.ToUpper(str))

	newStr:= "one,two,three,four"
	split:=strings.Split(newStr, ",")

	fmt.Println(split)
	fmt.Println(reflect.TypeOf(split))
}