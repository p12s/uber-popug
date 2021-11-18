package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/p12s/uber-popug/auth/pkg/models"
	"github.com/p12s/uber-popug/auth/pkg/repository"
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

// Authorization - signup/signin
type Authorization interface {
	CreateAccount(account models.Account) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
	GetAccountById(accountId int) (models.Account, error)
}

// AuthService - service
type AuthService struct {
	repo repository.Authorization
}

// NewAuthService - constructor
func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateAccount(account models.Account) (int, error) {
	account.Password = generatePasswordHash(account.Password)
	return s.repo.CreateAccount(account)
}

// GenerateToken - token generation
func (s *AuthService) GenerateToken(username, password string) (string, error) {
	account, err := s.repo.GetAccount(username, generatePasswordHash(password))
	if err != nil {
		return "", errors.New("account with this login/pass is not found")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		account.Id,
	})
	return token.SignedString([]byte(signingKey))
}

// ParseToken - getting authorized data from token
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

func (s *AuthService) GetAccountById(accountId int) (models.Account, error) {
	return s.repo.GetAccountById(accountId)
}
