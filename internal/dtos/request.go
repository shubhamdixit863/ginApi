package dtos

import (
	"ginTraining/internal/entity"
	"time"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string
	Age      int
}

func (sr SignupRequest) ToEntity() *entity.User {
	return &entity.User{
		Username:  sr.Username,
		Password:  sr.Password,
		Name:      sr.Name,
		Age:       sr.Age,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

}
