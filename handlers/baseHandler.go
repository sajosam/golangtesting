package handlers

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

)

type Handler struct {
	DB *gorm.DB
}

func (handler *Handler) Connect(host, user, password, dbname, port string) {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	handler.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
}