package usermngmnt

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

func HealthCheck(c *gin.Context) {
	c.Status(http.StatusOK)
	c.String(http.StatusOK, "Super Secret Area")
}

func (usrHandler *UsrHandler) GetUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var users []User
	usrHandler.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

func (usrHandler *UsrHandler) AddUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var user User
	json.NewDecoder(c.Request.Body).Decode(&user)
	usrHandler.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}

func (usrHandler *UsrHandler) GetUserInd(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var user User
	id := c.Param("id")
	usrHandler.DB.First(&user, id)
	c.JSON(http.StatusOK, user)
}

func (usrHandler *UsrHandler) DelUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var user User
	id := c.Param("id")
	usrHandler.DB.Delete(&user, id)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

func (usrHandler *UsrHandler) UpdateUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var user User
	id := c.Param("id")
	usrHandler.DB.First(&user, id)
	json.NewDecoder(c.Request.Body).Decode(&user)
	usrHandler.DB.Save(&user)
	c.JSON(http.StatusOK, user)
}
