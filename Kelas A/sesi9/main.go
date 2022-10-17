package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var url = "https://jsonplaceholder.typicode.com/posts"
var port = ":4000"

func main() {
	http.HandleFunc("/posts", AllowOnlyGet(Auth(GetPostById)))

	log.Println("server running at port", port)
	http.ListenAndServe(port, nil)
}

func GetPostById(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	resp, err := HttpGet(url + "/" + query.Get("id"))
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"payload": resp,
	})
}

func AllowOnlyGet(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "method not allowed!",
			})
			return
		}
		next(w, r)
	}
	// return method == http.MethodGet
}

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "need basic auth",
			})
			return
		}

		if username != "admin" || password != "admin" {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "username / password salah.",
			})
			return
		}

		next(w, r)
	}
}

type Post struct {
	Id     int
	UserId int
	Title  string
	Body   string
}

func get(url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	var data Post
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", data)
}

func HttpGet(url string) (interface{}, error) {
	data, err := req(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func HttpPost(url string) {

	reqPayload := Post{
		UserId: 1,
		Title:  "Ini title",
		Body:   "Ini body",
	}

	data, err := json.Marshal(reqPayload)
	if err != nil {
		panic(err)
	}

	resp, err := req(http.MethodPost, url, data)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}

func Update(url string) {
	reqPayload := Post{
		UserId: 1,
		Title:  "Ini title",
		Body:   "Ini body",
	}

	data, err := json.Marshal(reqPayload)
	if err != nil {
		panic(err)
	}

	resp, err := req(http.MethodPut, url, data)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}

func req(method, url string, body []byte) (map[string]interface{}, error) {
	client := http.Client{}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil

}
