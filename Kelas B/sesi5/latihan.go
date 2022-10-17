package main

import (
	"fmt"
	"math/rand"
	"time"
)

var breakPoint = 11

func main() {
	// Buatlah sebuah aplikasi permainan korek api bernyanyi
	// ada 4 player (4 goroutine)

	c := make(chan *korek)
	done := make(chan *korek)

	go player("User 1", c, done)
	go player("User 2", c, done)
	go player("User 3", c, done)
	go player("User 4", c, done)

	clear(c, done)
}

type korek struct {
	hits       int
	lastPlayer string
}

func player(name string, c, done chan *korek) {

	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 100
	for {
		select {
		case data := <-c:
			v := rand.Intn(max-min) + min
			fmt.Println(name, v, breakPoint)
			time.Sleep(500 * time.Millisecond)

			data.hits++
			data.lastPlayer = name
			if v%breakPoint == 0 {
				fmt.Println("Korek berhenti di", name)
				done <- data
				return
			}
			c <- data

			fmt.Println("korek ada di", name, "pada hits ke", data.hits)
		}
	}
}

func clear(c, done chan *korek) {
	c <- new(korek)
	for {
		select {
		case d := <-done:
			fmt.Println("Yg kalah adalah", d.lastPlayer)
			return
		}
	}
}
