package main

import (
	"fmt"
)

func hello() {
	defer fmt.Println("Ini defer dari Hello")
	fmt.Println("Hello 2")
	panic("error")
}

func main() {
	// defer fmt.Println("Defer 1")
	// defer fmt.Println("Defer 4")
	hello()
	fmt.Println("Hello 1")
	// num := 10
	// switch {
	// case num > 9:
	// 	fmt.Println("A")
	// 	return
	// case num > 8:
	// 	fmt.Println("AB")
	// 	return
	// case num > 7:
	// 	fmt.Println("B")
	// 	return
	// case num > 6:
	// 	fmt.Println("BC")
	// 	return
	// }
}
