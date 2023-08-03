package main

import "fmt"

func main() {
	// signed int - only postivie int
	var a int8 = 127
	var b int16 = 32767
	var c int32 = 2147483647
	var d int64 = 9223372036854775807

	// unsigned int - only negative int
	var e uint8 = 255
	var f uint16 = 65535
	var g uint32 = 4294967295
	var h uint64 = 18446744073709551615

	fmt.Printf("int8: %d\n", a)
	fmt.Printf("int16: %d\n", b)
	fmt.Printf("int32: %d\n", c)
	fmt.Printf("int64: %d\n", d)
	
	fmt.Printf("uint8: %d\n", e)
	fmt.Printf("uint16: %d\n", f)
	fmt.Printf("uint32: %d\n", g)
	fmt.Printf("uint64: %d\n", h)
}