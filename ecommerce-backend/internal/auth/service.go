package auth

import (
	"ecommerce-backend/internal/domain"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(user *domain.User) error
	Login(email, password string) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type service struct {
	repo       Repository
	jwtSecret  string
}

func NewService(repo Repository, jwtSecret string) Service {
	return &service{repo, jwtSecret}
}

func (s *service) Register(user *domain.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.repo.CreateUser(user)
}

func (s *service) Login(email, password string) (string, error) {
	user, err := s.repo.FindUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.jwtSecret))
}

func (s *service) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.jwtSecret), nil
	})
}