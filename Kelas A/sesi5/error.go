package main

import (
	"fmt"
)

type req struct {
	name string
	age  int
}

func main() {
	defer catch()
	r := req{
		name: "Hacktiv8",
		age:  10,
	}

	err := validate(&r)
	if err != nil {
		panic(err)
		// return
	}

	process(&r)

	fmt.Println(r.name)
}

func process(r *req) {
	// process
	if r.age <= 13 {
		panic("error")
	}
}

func catch() {
	if r := recover(); r != nil {
		fmt.Println("error :", r)
	}
}

func validate(r *req) error {
	if r.name == "" {
		return fmt.Errorf("Name tidak boleh kosong")
	}

	if len(r.name) <= 4 {
		return fmt.Errorf("Name harus lebih besar dari 4")
	}

	if r.age <= 13 {
		return fmt.Errorf("Umur harus lebih besar dari 13, got - %d", r.age)
	}
	return nil
}
