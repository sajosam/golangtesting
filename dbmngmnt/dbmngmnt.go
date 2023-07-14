package dbmngmnt

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        int    `json:"ID"`
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
}

type Book struct {
	gorm.Model
	ID          int    `json:"ID"`
	BookName    string `json:"book_name"`
	Price       float64 `json:"price"`
}
// type UsrHandler struct {
// 	DB *gorm.DB
// }

