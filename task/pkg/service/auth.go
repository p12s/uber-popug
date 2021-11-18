package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/p12s/uber-popug/task/pkg/models"
	"github.com/p12s/uber-popug/task/pkg/repository"
)

const (
	salt       = "8284kjalsdf282-4asfjae93sdf"
	tokenTTL   = 12 * time.Hour
	signingKey = "29dsjkadf*^(&le23#ls93s02a0d9"
)

// tokenClaims - tooken object
type tokenClaims struct {
	jwt.StandardClaims
	AccountId int `json:"account_id"`
}

type Authorizer interface {
	ParseToken(token string) (int, error)
	GetAccount(token string) (models.Account, error)
	GetAccountById(publicId uuid.UUID) (models.Account, error)
}

// AuthService - service
type AuthService struct {
	repo repository.Authorizer
}

// NewAuthService - constructor
func NewAuthService(repo repository.Authorizer) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) GetAccount(token string) (models.Account, error) {
	return s.repo.GetAccount(token)
}

func (s *AuthService) GetAccountById(publicId uuid.UUID) (models.Account, error) {
	return s.repo.GetAccountById(publicId)
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.AccountId, nil
}

// generatePasswordHash - hash generare from password
func generatePasswordHash(password string) string {
	hash := sha1.New() // #nosec
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
