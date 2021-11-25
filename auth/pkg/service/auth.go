package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
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

type Authorizer interface {
	CreateAccount(account models.Account) (int, error)
	UpdateAccount(input models.UpdateAccountInput) error
	DeleteAccountByPublicId(accountPublicId uuid.UUID) error
	GenerateToken(username, password string) (models.AccountToken, error)
	ParseToken(token string) (int, error)
	GetAccountById(accountId int) (models.Account, error)
}

// AuthService - service
type AuthService struct {
	repo repository.Authorizer
}

// NewAuthService - constructor
func NewAuthService(repo repository.Authorizer) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateAccount(account models.Account) (int, error) {
	account.Password = generatePasswordHash(account.Password)
	return s.repo.CreateAccount(account)
}

func (s *AuthService) UpdateAccount(input models.UpdateAccountInput) error {
	if input.Password != nil {
		*input.Password = generatePasswordHash(*input.Password)
	}
	return s.repo.UpdateAccount(input)
}

func (s *AuthService) DeleteAccountByPublicId(accountPublicId uuid.UUID) error {
	return s.repo.DeleteAccountByPublicId(accountPublicId)
}

func (s *AuthService) GenerateToken(username, password string) (models.AccountToken, error) {
	var accountToken models.AccountToken
	account, err := s.repo.GetAccount(username, generatePasswordHash(password))
	if err != nil {
		return accountToken, errors.New("account with this login/pass is not found")
	}
	accountToken.PublicId = account.PublicId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		account.Id,
	})
	accountToken.Token, err = token.SignedString([]byte(signingKey))
	if err != nil {
		return accountToken, fmt.Errorf("token generate error: %x", err.Error())
	}

	return accountToken, nil
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

func (s *AuthService) GetAccountById(accountId int) (models.Account, error) {
	return s.repo.GetAccountById(accountId)
}
