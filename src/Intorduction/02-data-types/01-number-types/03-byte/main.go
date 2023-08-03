package main

import "fmt"

func main() {
	var ch byte = 90
	fmt.Printf("%08b %02x %v\n", ch, ch, ch)

	var ch0 byte = 'Z'
	fmt.Printf("%08b %02x %v\n", ch0, ch0, ch0)

	var ch1 byte = 0b01011010
	var ch2 byte = 0o132
	var ch3 byte = 0x5a

	fmt.Printf("%08b %02x %v\n", ch1, ch1, ch1)
	fmt.Printf("%08b %02x %v\n", ch2, ch2, ch1)
	fmt.Printf("%08b %02x %v\n", ch3, ch3, ch1)
}