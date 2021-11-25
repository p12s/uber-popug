package service

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/p12s/uber-popug/analitycs/pkg/models"
	"github.com/p12s/uber-popug/analitycs/pkg/repository"
)

const (
	salt       = "8284kjalsdf282-4asfjae93sdf"
	tokenTTL   = 12 * time.Hour
	signingKey = "29dsjkadf*^(&le23#ls93s02a0d9"
)

type tokenClaims struct {
	jwt.StandardClaims
	AccountId int `json:"account_id"`
}

type Authorizer interface {
	CreateAccount(account models.Account) (int, error)
	UpdateAccount(input models.UpdateAccountInput) error
	DeleteAccountByPublicId(accountPublicId uuid.UUID) error
	ParseToken(token string) (int, error)
	// это костыльный метод - опираться нужно только на publicId uuid.UUID
	GetAccountByPrimaryId(accountId int) (models.Account, error)
}

type AuthService struct {
	repo repository.Authorizer
}

func NewAuthService(repo repository.Authorizer) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateAccount(account models.Account) (int, error) {
	return s.repo.CreateAccount(account)
}

func (s *AuthService) UpdateAccount(input models.UpdateAccountInput) error {
	return s.repo.UpdateAccount(input)
}

func (s *AuthService) DeleteAccountByPublicId(accountPublicId uuid.UUID) error {
	return s.repo.DeleteAccountByPublicId(accountPublicId)
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

func (s *AuthService) GetAccountByPrimaryId(accountId int) (models.Account, error) {
	return s.repo.GetAccountByPrimaryId(accountId)
}
