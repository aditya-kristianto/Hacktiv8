package main

import (
	"fmt"
)

func introduce(name string, c chan<- string) {
	text := fmt.Sprintf("Hello %s", name)

	c <- text
	// fmt.Println()
}

func print(c <-chan string) {
	fmt.Println(<-c)
}

func calculate(nums []int, c chan int) {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	// fmt.Println(sum)
	c <- sum
}

func main() {
	c := make(chan int)
	nums := []int{1, 2, 3, 4, 5, 6}
	// calculate(nums, c)
	go calculate(nums[:len(nums)/2], c)
	go calculate(nums[len(nums)/2:], c)

	n1 := <-c
	n2 := <-c
	fmt.Println(n1 + n2)
}
