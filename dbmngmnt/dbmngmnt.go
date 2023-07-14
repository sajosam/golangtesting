package dbmngmnt

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        int    `json:"ID"`
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
}

type UsrHandler struct {
	DB *gorm.DB
}

func (usrHandler *UsrHandler) Connection(host, user, password, dbname, port string) {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	usrHandler.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	usrHandler.DB.AutoMigrate(&User{})
}