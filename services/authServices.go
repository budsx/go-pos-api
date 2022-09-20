package services

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

type AuthService interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (bool, int, error)
}

type JwtService struct {
	UserID int
	jwt.RegisteredClaims
}

var SECRET_KEY = []byte(os.Getenv("SECRET_KEY"))

func NewService() *JwtService {
	return &JwtService{}
}

func (s *JwtService) GenerateToken(userID int) (string, error) {
	claim := JwtService{}
	claim.UserID = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *JwtService) ValidateToken(encodedToken string) (bool, int, error) {
	token, err := jwt.ParseWithClaims(encodedToken, &JwtService{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return false, 0, err
	}

	if claim, ok := token.Claims.(*JwtService); ok && token.Valid {
		return true, claim.UserID, nil
	} else {
		return false, claim.UserID, err
	}
}
