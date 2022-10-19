package user

import (
	"errors"
	"final-project/internal/app/model"
	"final-project/internal/app/user/repository"
	"final-project/internal/app/user/repository/postgres"
	"final-project/internal/pkg/helper"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
	Service struct {
		repository repository.Repository
	}
)

func NewService(db *gorm.DB) *Service {
	return &Service{
		repository: postgres.NewRepository(db),
	}
}

func (s *Service) Register(req *RegisterRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user, err := s.repository.Read(req.Email)
	if err != nil {
		return err
	}
	if user.Username == req.Username {
		return errors.New("username already registered")
	} else if user.Email == req.Email {
		return errors.New("email already registered")
	}

	data := &model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
		Age:      req.Age,
	}
	_, err = s.repository.Create(data)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Login(req *LoginRequest) (string, error) {
	user, err := s.repository.Read(req.Email)
	if err != nil {
		return "", err
	}

	if user.Email == "" {
		return "", errors.New("email is not registered")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", err
	}

	claims := &helper.JwtCustomClaims{
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (s *Service) UpdateUser(userID *uuid.UUID, data *model.User) (*model.User, error) {
	user, err := s.repository.ReadByID(userID)
	if err != nil {
		return nil, err
	}

	if user.ID == uuid.Nil {
		return nil, errors.New("Unauthorized")
	}

	user, err = s.repository.Update(userID, data)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) DeleteUser(userID *uuid.UUID) error {
	_, err := s.repository.Delete(userID)
	if err != nil {
		return err
	}

	return nil
}
