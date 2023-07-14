package bookmngmnt

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/testapi/dbmngmnt"
)



type BookHandler struct {
	DB *gorm.DB
}

func (bookHandler *BookHandler) Connection(host, user, password, dbname, port string) {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	bookHandler.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	bookHandler.DB.AutoMigrate(&dbmngmnt.Book{})
}


func HealthCheck(c *gin.Context) {
	c.Status(http.StatusOK)
	c.String(http.StatusOK, "Super Secret Area")
}

func (bookHandler *BookHandler) GetBook(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var book []dbmngmnt.Book
	bookHandler.DB.Find(&book)
	c.JSON(http.StatusOK, book)
}

func (bookHandler *BookHandler) AddBook(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var book dbmngmnt.Book
	json.NewDecoder(c.Request.Body).Decode(&book)
	bookHandler.DB.Create(&book)
	c.JSON(http.StatusCreated, book)
}

func (bookHandler *BookHandler) GetBookInd(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var book dbmngmnt.Book
	id := c.Param("id")
	bookHandler.DB.First(&book, id)
	c.JSON(http.StatusOK, book)
}

func (bookHandler *BookHandler) DelBook(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var book dbmngmnt.Book
	id := c.Param("id")
	bookHandler.DB.Delete(&book, id)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}

func (bookHandler *BookHandler) UpdateBook(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var book dbmngmnt.Book
	id := c.Param("id")
	bookHandler.DB.First(&book, id)
	json.NewDecoder(c.Request.Body).Decode(&book)
	bookHandler.DB.Save(&book)
	c.JSON(http.StatusOK, book)
}
