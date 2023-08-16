package authentication

import (
	"os"
	"time"

	"erp-system.com/v1/domain/model"
	"github.com/dgrijalva/jwt-go"
)

var (
	secretKey = []byte(os.Getenv("SECRET_JWT"))
)

type CustomClaims struct {
	User model.PersonJWT `json:"person"`
	jwt.StandardClaims
}

func CreateToken(p model.Person) (string, error) {
	claims := CustomClaims{
		model.PersonJWT{
			Name:     p.Name,
			ID:       p.ID,
			Email:    p.Email,
			IsMaster: p.IsMaster,
		},
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			IssuedAt:  time.Now().Unix(),
			Subject:   "auth",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
