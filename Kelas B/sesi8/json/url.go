package main

import (
	"fmt"
	"net/url"
)

func main() {
	rawUrl := "https://www.noobee.id/search?title=&categories=&types=&tags=&kind=Artikel,Course"

	url, err := url.Parse(rawUrl)
	if err != nil {
		panic(err)
	}

	fmt.Println("raw url", rawUrl)
	fmt.Println("Scheme", url.Scheme)
	fmt.Println("Host", url.Host)
	fmt.Println("Path", url.Path)
	fmt.Println("Port", url.Port())

	query := url.Query()
	fmt.Println("Query", query)
	fmt.Println("Query Kind :", query.Get("kind"))
	fmt.Println("Query Kind :", query["kind"])
}
