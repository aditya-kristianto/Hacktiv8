package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("%d genap\n", i)
		} else {
			fmt.Printf("%d ganjil\n", i)
		}
	}
}
