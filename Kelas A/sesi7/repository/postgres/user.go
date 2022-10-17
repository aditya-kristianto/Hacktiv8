package postgres

import (
	"database/sql"
	"sesi7/model"
	"sesi7/repository"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) repository.UserRepository {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) CreateUser(user *model.User) error {
	query := `
		INSERT INTO users (name, email)
		VALUES ($1, $2)
	`

	stmt, err := u.db.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email)
	if err != nil {
		return err
	}

	return nil
}

func (u *userRepo) GetUsers() (*[]model.User, error) {
	query := `
		SELECT id, name, email, created_at
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

	var users []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(
			&user.ID, &user.Name, &user.Email, &user.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return &users, nil
}
