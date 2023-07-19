package services

import (
	"errors"
	"fmt"
	"ginTraining/internal/dtos"
	"ginTraining/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
	"time"
)

type AuthService struct {
	// Pass the dependecy for the repository
	AuthRepo repository.IAuthRepository
}

func (as *AuthService) SignupUser(signupRequest dtos.SignupRequest) (string, error) {

	// We will make a conversion here from dto to entity

	_, err := as.AuthRepo.Signup(signupRequest.ToEntity())
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	/// We will just create jwt and all ---->for the user signup

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": signupRequest.Username,
		"nbf":      time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		"exp":      time.Now().Add(10 * time.Minute),
	})

	log.Println(os.Getenv("JWT_SECRET"))

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (as *AuthService) LoginUser(loginRequest dtos.LoginRequest) (string, error) {

	login, err := as.AuthRepo.Login(loginRequest.Username, loginRequest.Password)
	if err != nil {
		return "", err
	}

	// we have to compare the passwords here

	if login.Password != loginRequest.Password {
		return "", errors.New("Passwords donot match")
	}

	// Send jwt token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": loginRequest.Username,
		"nbf":      time.Date(2023, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		"exp":      time.Now().Add(10 * time.Minute),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil

}
