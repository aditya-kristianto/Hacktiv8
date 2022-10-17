package main

import (
	"log"
	"net/http"
	"sesi9/controller"
	"sesi9/httpclient"
)

var url = "https://jsonplaceholder.typicode.com/posts"
var port = ":4000"

func main() {
	client := httpclient.NewClient(10, url)

	postController := controller.NewPostController(client)

	http.HandleFunc("/posts", postController.GetPosts)

	log.Println("server running at port", port)
	http.ListenAndServe(port, nil)
}
