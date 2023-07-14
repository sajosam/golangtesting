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
	var users []dbmngmnt.User
	bookHandler.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

func (bookHandler *BookHandler) AddBook(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var user dbmngmnt.User
	json.NewDecoder(c.Request.Body).Decode(&user)
	bookHandler.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}

func (bookHandler *BookHandler) GetBookInd(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var user dbmngmnt.User
	id := c.Param("id")
	bookHandler.DB.First(&user, id)
	c.JSON(http.StatusOK, user)
}

func (bookHandler *BookHandler) DelBook(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var user dbmngmnt.User
	id := c.Param("id")
	bookHandler.DB.Delete(&user, id)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

func (bookHandler *BookHandler) UpdateBook(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var user dbmngmnt.User
	id := c.Param("id")
	bookHandler.DB.First(&user, id)
	json.NewDecoder(c.Request.Body).Decode(&user)
	bookHandler.DB.Save(&user)
	c.JSON(http.StatusOK, user)
}
