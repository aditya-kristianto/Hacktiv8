package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	IdUser   int `json:"idUser"`
	Email    string
	Password string
	Role     string
	Name     string
	Age      int
}

func main() {
	userString := `
		[
			{
				"idUser" : 1,
				"email" : "kelas.b@mnc.com",
				"password" : "mncb",
				"role" : "member",
				"full_name" : "MNC B",
				"age" : 10
			}
		]
	`

	var user []User
	err := json.Unmarshal([]byte(userString), &user)
	if err != nil {
		panic(err)
	}

	var userMap = []map[string]interface{}{}
	err = json.Unmarshal([]byte(userString), &userMap)
	if err != nil {
		panic(err)
	}

	var userInterface interface{}
	err = json.Unmarshal([]byte(userString), &userInterface)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", user[0].Email)
	fmt.Printf("%+v\n", userMap[0]["email"])

	userStructByte, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	userMapByte, err := json.Marshal(userMap)
	if err != nil {
		panic(err)
	}
	userInterfaceByte, err := json.Marshal(userInterface)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", string(userStructByte))
	fmt.Printf("%+v\n", string(userMapByte))
	fmt.Printf("%+v\n", string(userInterfaceByte))

}
