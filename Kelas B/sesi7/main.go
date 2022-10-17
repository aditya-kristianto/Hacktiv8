package main

import (
	"fmt"
	"math/rand"
	"sesi7/config"
	"sesi7/gin/controller"
	"sesi7/gin/repository/gorm"
	"sesi7/gin/router"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	err := config.ConnectGorm()
	if err != nil {
		panic(err)
	}
	err = config.ConnectPostgres()
	if err != nil {
		panic(err)
	}

	// pos := config.GetPostgres()
	orm := config.GetGORM()
	empRepo := gorm.NewUserRepo(orm)

	employeeHandler := controller.NewEmployeeHandler(empRepo)

	router := router.NewRouter(employeeHandler)
	router.CreateRouter().Run(":4000")

	// CreateUser(2)
	// GetUsers()
	// GetUserByID(12)
}

type User struct {
	ID        int
	Name      string
	Email     string
	Contact   string
	CreatedAt time.Time
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func CreateUser(len int) {
	db := config.GetPostgres()
	var users []User
	for i := 0; i < len; i++ {
		users = append(users, User{
			Name:      RandStringRunes(10),
			Email:     fmt.Sprintf("%s@gmail.com", RandStringRunes(20)),
			Contact:   RandStringRunes(10),
			CreatedAt: time.Now(),
		})
	}

	query := `
		INSERT INTO users (
			name, email, contact, created_at
		) VALUES ($1, $2, $3, $4)
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	for _, u := range users {
		_, err := stmt.Exec(u.Name, u.Email, u.Contact, u.CreatedAt)
		if err != nil {
			fmt.Println("error :", err.Error())
		}
	}

	fmt.Println("created", len, "users")
}

func GetUsers() {
	db := config.GetPostgres()

	query := `
		SELECT 
			id, name, email, contact, created_at
		FROM users
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	var users []User
	rows.Next()
	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.ID, &user.Name, &user.Email, &user.Contact, &user.CreatedAt,
		)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}

	for _, u := range users {
		fmt.Println("name :", u.Name)
	}

}

func GetUserByID(id int) {
	db := config.GetPostgres()

	query := `
		SELECT 
			id, name, email, contact, created_at
		FROM users
		WHERE id=$1
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)

	var user User
	err = row.Scan(
		&user.ID, &user.Name, &user.Email, &user.Contact, &user.CreatedAt,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)

}
