package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"full_name"`
	Address string `json:"address"`
}

func main() {
	strData := `
		[
			{
				"id" : 1,
				"full_name" : "MNC A",
				"address" : "Jakarta"
			},
			{
				"id" : 2,
				"full_name" : "MNC B",
				"address" : "Jakarta"
			}
		]
	`

	var user []User
	// var user []map[string]interface{} = []map[string]interface{}{}
	err := json.Unmarshal([]byte(strData), &user)
	if err != nil {
		panic(err)
	}

	fmt.Println(user)

	user1 := []User{
		{
			Id:      1,
			Name:    "MNC B",
			Address: "Jakarta",
		},
		{
			Id:      2,
			Name:    "MNC A",
			Address: "Jakarta",
		},
	}

	user1Byte, err := json.Marshal(user1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(user1Byte))
}
