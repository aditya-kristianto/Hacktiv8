package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	defer catch()
	req := "10a"
	err := validate(req)
	if err != nil {
		panic(err)
	}
	processData(req)

}

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Ada error nih ...", r)
	} else {
		fmt.Println("Aman..")
	}
}

func validate(d string) error {
	_, err := strconv.Atoi(d)
	if err != nil {
		msg := fmt.Sprintf("Error : %s", err.Error())
		return errors.New(msg)
	}

	return nil
}

func processData(d string) {
	data, _ := strconv.Atoi(d)
	data += 10

	fmt.Println(data)
}
