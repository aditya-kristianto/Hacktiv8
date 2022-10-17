package main

import (
	"fmt"
	"time"
)

type Text struct {
	msg string
	err error
}

func main() {
	c := make(chan string)
	c2 := make(chan string)
	c3 := make(chan *Text)

	go func() {
		time.Sleep(3 * time.Second)
		c <- "Ini dari c1"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "Ini dari c2"
	}()

	go introduce("Kami", c3)

	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case v := <-c2:
			fmt.Println(v)
		case v := <-c3:
			if v.err != nil {
				fmt.Println("Ada error nih", v.err)
			} else {
				fmt.Println(v.msg)

			}

		case <-time.After(2 * time.Second):
			fmt.Println("TIMEOUT ...")
			return
		}
	}
}

func print(c <-chan string) {
	fmt.Println(<-c)
}

func introduce(name string, c chan *Text) {
	text := fmt.Sprintf("Hello %s", name)
	if len(name) == 0 {
		c <- &Text{
			err: fmt.Errorf("ga boleh kosong"),
		}
		return
	}

	c <- &Text{
		msg: text,
		err: nil,
	}
}
