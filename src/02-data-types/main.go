package main

import "fmt"

func main() {
	// format
	str:="hello"
	num1:=10
	num2:=3.14159
	num3:=1234567890
	boolean:=true
	char:='A'

	// string formatter
	fmt.Printf("%T\n", str) // type of var
	fmt.Printf("%v\n", str) // default format
	fmt.Printf("%s\n", str) // print string
	fmt.Printf("%q\n", str) // print quoted string
	fmt.Printf("%x\n", []byte(str)) // print hex endocing of bytes
	fmt.Printf("%X\n", []byte(str)) // print uppercase hex endocing of bytes
	fmt.Println()

	// int formatter
	fmt.Printf("%d\n", num1) // print decimal int
	fmt.Printf("%b\n", num1) // print binay int
	fmt.Printf("%o\n", num1) // print octal int
	fmt.Printf("%x\n", num1) // print hex encoding int
	fmt.Printf("%X\n", num1) // print upppercase hex encoding int
	fmt.Printf("%c\n", char) // print char
	fmt.Println()

	// floating point
	fmt.Printf("%f\n", num2) // print floating point num
	fmt.Printf("%e\n", num2) // print floating point in scientfic natation
	fmt.Printf("%E\n", num2) // print floating point in scientfic natation uppercase
	fmt.Printf("%g\n", num2) // print floating-point number in decimal or scientific notation, depending on the value
	fmt.Printf("%G\n", num2) // print floating-point number in decimal or scientific notation, depending on the value uppercase
	fmt.Println()

	// width and precision
	fmt.Printf("|%5d|\n", num1)
	fmt.Printf("|%-5d|\n", num1)
	fmt.Printf("|%5.2f|\n", num2)
	fmt.Printf("|%-5.2f|\n", num2)
	fmt.Println()

	// boolean
	fmt.Printf("%t\n", boolean)

	// pointer
	fmt.Printf("%p\n", &num3)
}