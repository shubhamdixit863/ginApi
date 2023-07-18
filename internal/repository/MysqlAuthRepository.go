package repository

import (
	"ginTraining/internal/entity"
	"gorm.io/gorm"
)

type MysqlAuthRepository struct {

	// We will define our database object
	Db *gorm.DB
}

func (ma *MysqlAuthRepository) Signup(user *entity.User) (int64, error) {

	record := ma.Db.Create(user)
	if record.Error != nil {
		return 0, record.Error
	}

	return record.RowsAffected, nil

}
func (ma *MysqlAuthRepository) Login(username, password string) (*entity.User, error) {

	user := entity.User{
		Username: username,
	}
	data := ma.Db.First(&user)
	if data.Error != nil {
		return nil, data.Error
	}

	return &user, nil
}
