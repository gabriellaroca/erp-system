package api

import (
	"net/http"

	authentication "erp-system.com/v1/domain/Security/Authentication/jwt"
	"erp-system.com/v1/domain/usecase"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AuthHandler struct {
	personUsecase usecase.PersonUsecase
}

func NewAuthHandler(personUsecase usecase.PersonUsecase) *AuthHandler {
	return &AuthHandler{personUsecase: personUsecase}
}

func (handler *AuthHandler) Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	person, err := handler.personUsecase.ReadByEmail(email)
	if err != nil {
		logrus.Errorf("Error reading user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An internal error occurred"})
		return
	}

	if person.Password == password {
		token, err := authentication.CreateToken(person)
		if err != nil {
			logrus.Errorf("Error creating token: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "An internal error occurred"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
}
