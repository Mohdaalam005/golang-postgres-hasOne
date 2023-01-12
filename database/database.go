package database

import (
	"github.com/mohdaalam005/one-to-one/helpers"
	"github.com/mohdaalam005/one-to-one/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var Db *gorm.DB

func Connect() {

	url := os.Getenv("url")
	var err error
	Db, err = gorm.Open(postgres.Open(url), &gorm.Config{})
	helpers.CheckNilError(err)
	Db.AutoMigrate(&models.User{}, &models.CreditCard{})

}
