package repository

import (
	"fmt"
	"ginTraining/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"testing"
	"time"
)

func MysqlConnect(uri string) *gorm.DB {

	db, err := gorm.Open(mysql.Open(uri), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Connected with Mysql")

	return db
}

func TestConnection(t *testing.T) {

	// We will test the connection
	dsn := "admin:o9Uusjfn@tcp(mysql-135552-0.cloudclusters.net:17741)/mysqlGolang?charset=utf8mb4&parseTime=True&loc=Local"
	db := MysqlConnect(dsn)
	assert.NotNil(t, db)

}

func TestInitDb(t *testing.T) {

	dsn := "admin:o9Uusjfn@tcp(mysql-135552-0.cloudclusters.net:17741)/mysqlGolang?charset=utf8mb4&parseTime=True&loc=Local"

	ma := &MysqlAuthRepository{
		Db: MysqlConnect(dsn),
	}

	err := ma.Db.AutoMigrate(&entity.User{})
	assert.Nil(t, err)

}

func TestMysqlAuthRepository_Signup(t *testing.T) {
	dsn := "admin:o9Uusjfn@tcp(mysql-135552-0.cloudclusters.net:17741)/mysqlGolang?charset=utf8mb4&parseTime=True&loc=Local"

	ma := &MysqlAuthRepository{
		Db: MysqlConnect(dsn),
	}

	data, err := ma.Signup(&entity.User{
		Username:  "jon123",
		Password:  "123456",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	assert.Nil(t, err)
	assert.Greater(t, 0, data)

}
