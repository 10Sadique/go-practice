package main

import "fmt"

func main() {
	var map1 map[int]int
	if map1 == nil {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	map2 := map[int]string{
		90: "Dog",
		91: "Cat",
		92: "Cow",
		93: "Bird",
		94: "Rabbit",
	}

	fmt.Println("Map-2: ", map2)
	fmt.Println(map2[90])
	fmt.Println(len(map2))

	delete(map2, 91)
	fmt.Println("Map-2: ", map2)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("Map", n)
}
