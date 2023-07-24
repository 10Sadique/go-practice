package main

import "fmt"

func main() {
	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set", s)
	fmt.Println("get", s[2])
	fmt.Println("len", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apnd: ", s)

	c := make([]string, len(s))
	copy(c, s)

	fmt.Println(c)

	copySlice := s[2:5]
	fmt.Println("s1", copySlice)

	copySlice = s[:5]
	fmt.Println("s2", copySlice)

	towD := make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerLen := i + 1
		towD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			towD[i][j] = i + j
		}
	}

	fmt.Println(towD)
}
