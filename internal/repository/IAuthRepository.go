package repository

import "ginTraining/internal/entity"

type IAuthRepository interface {
	Signup(user *entity.User) (int64, error)
	Login(username, password string) (*entity.User, error)
}
