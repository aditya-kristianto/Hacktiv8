package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("World")
	go deferFunc()
	time.Sleep(1 * time.Second)
	nilai := 10

	switch {
	case nilai > 9:
		fmt.Println("Bagoos")
		return
	case nilai > 8:
		fmt.Println("Okee lah")
		return
	default:
		fmt.Println("Ayo belajar lagi")
		return
	}
}

func deferFunc() {
	defer fmt.Println("ini kena defer")
	defer fmt.Println("ini kena defer 3")
	defer fmt.Println("ini kena defer 2")
	fmt.Println("Ini dari defer function")
	os.Exit(1)
}

/*
1. 26,24,25,7,12
2. 26,25,24,7,12

*/
