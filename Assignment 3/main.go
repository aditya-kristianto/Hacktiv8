package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/robfig/cron/v3"
)

// Data
type Data struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

// Response
type Response struct {
	Status Data `json:"status"`
}

func main() {
	c := cron.New()
	c.AddFunc("@every 15s", setJSON)
	c.Start()

	e := echo.New()
	e.Static("/assets", "assets")
	e.GET("/json", getJSON)
	e.File("/", "public/index.html")
	e.Logger.Fatal(e.Start(":1323"))
}

func setJSON() {
	directory := "assets"
	filename := "status.json"
	if _, err := os.Stat(directory + "/" + filename); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(directory, 0755)
		if err != nil {
			fmt.Println(err)
		}

		_, err = os.Create(directory + "/" + filename)
		if err != nil {
			fmt.Println(err)
		}
	}

	response := &Response{
		Status: Data{
			Water: getRandomValue(),
			Wind:  getRandomValue(),
		},
	}

	content, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile(directory+"/"+filename, content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func getJSON(c echo.Context) error {
	response := &Response{
		Status: Data{
			Water: getRandomValue(),
			Wind:  getRandomValue(),
		},
	}

	return c.JSON(http.StatusOK, response)
}

func getRandomValue() int {
	min := 1
	max := 100

	return rand.Intn(max-min) + min
}
