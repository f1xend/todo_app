package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	// "github.com/dgrijalva/jwt-go"
	"github.com/f1xend/todo-app"
	"github.com/f1xend/todo-app/package/repository"
	"github.com/golang-jwt/jwt/v5"
)

const (
	salt      = "hjqrhjqw124617ajfhajs"
	tokenTTL  = 12 * time.Hour
	secretKey = "my_new_secret_key"
)

type TokenClaims struct {
	UserId int `json:"user_id"`
	jwt.RegisteredClaims
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	// get user from DB
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	claims := TokenClaims{
		user.Id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(secretKey))

	return ss, err
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*TokenClaims); !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	} else {
		return claims.UserId, nil
	}
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
