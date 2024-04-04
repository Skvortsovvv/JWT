package service

import (
	"JWT/pkg/config"
	"JWT/pkg/model"
	"JWT/pkg/store"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"time"
)

type Service struct {
	storage store.Store
}

func NewService(storage store.Store) *Service {
	return &Service{
		storage: storage,
	}
}

type tokenClaims struct {
	jwt.StandardClaims
	Id string `json:"id"`
}

func (s *Service) RegisterUser(login, password string) error {
	user := model.NewUser(uuid.New(), login, password, 1)
	err := s.storage.RegisterUser(user)
	if err != nil {
		return errors.Wrap(err, "s.storage.RegisterUser")
	}
	return nil
}

func (s *Service) GetUserToken(login, password string) (string, error) {
	id, err := s.storage.GetUser(login, config.GenerateHash(password))
	if err != nil {
		return "", errors.Wrap(err, "s.storage.GetUser")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(config.TokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		id.String(),
	})

	return token.SignedString([]byte(config.SignKey))
}

func (s *Service) ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(config.SignKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}

	return claims.Id, nil
}
