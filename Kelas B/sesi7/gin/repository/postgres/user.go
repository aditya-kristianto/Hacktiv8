package postgres

import (
	"database/sql"
	"sesi7/gin/repository"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) repository.UserRepo {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) CreateUser(user *repository.User) error {

	query := `
		INSERT INTO users (
			name, email, contact
		) VALUES ($1, $2, $3)
	`

	stmt, err := u.db.Prepare(query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, user.Contact)
	return err
}
func (u *userRepo) GetUsers() (*[]repository.User, error) {

	query := `
		SELECT 
			id, name, email, contact, created_at
		FROM users
	`

	stmt, err := u.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var users []repository.User
	for rows.Next() {
		var user repository.User
		err := rows.Scan(
			&user.ID, &user.Name, &user.Email, &user.Contact, &user.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return &users, nil
}
func (u *userRepo) GetUserByID(id int) (*repository.User, error) {
	return nil, nil
}
