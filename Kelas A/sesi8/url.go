package main

import (
	"fmt"
	"net/url"
)

func main() {
	urlString := "https://www.noobee.id/search?title=&categories=article,course&types=&tags=&kind=Course"

	url, err := url.Parse(urlString)
	if err != nil {
		panic(err)
	}

	fmt.Println("Protokol :", url.Scheme)
	fmt.Println("Path :", url.Path)
	fmt.Println("Host :", url.Host)

	query := url.Query()
	fmt.Println("categories :", query.Get("categories"))
	fmt.Println("categories :", query["categories"])
}
