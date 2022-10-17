package controller

import (
	"encoding/json"
	"net/http"
	"sesi9/httpclient"
)

type PostController struct {
	client *httpclient.Client
}

func NewPostController(client *httpclient.Client) *PostController {
	return &PostController{
		client: client,
	}
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}

type Post struct {
	Id     int    `json:"id"`
	UserId int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func (p *PostController) GetPosts(w http.ResponseWriter, r *http.Request) {
	resp, err := p.client.Get("/")
	if err != nil {
		WriteJsonResponse(w, &Response{
			Status:  http.StatusInternalServerError,
			Message: "INTERNAL_SERVER_ERROR",
			Error:   err.Error(),
		})
		return
	}

	var posts []Post
	err = json.Unmarshal(resp, &posts)
	if err != nil {
		WriteJsonResponse(w, &Response{
			Status:  http.StatusInternalServerError,
			Message: "INTERNAL_SERVER_ERROR",
			Error:   err.Error(),
		})
		return
	}
	WriteJsonResponse(w, &Response{
		Status:  http.StatusOK,
		Message: "GET_POSTS_SUCCESS",
		Payload: posts,
	})

}
