package bookmngmnt

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/testapi/models"
)


func HealthCheck(c *gin.Context) {
	c.Status(http.StatusOK)
	c.String(http.StatusOK, "Super Secret Area")
}

func (handler  *Handler) GetBook(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var book []models.Book
	bookHandler.DB.Find(&book)
	c.JSON(http.StatusOK, book)
}

func (handler  *Handler) AddBook(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var book models.Book
	json.NewDecoder(c.Request.Body).Decode(&book)
	bookHandler.DB.Create(&book)
	c.JSON(http.StatusCreated, book)
}

func (handler  *Handler) GetBookInd(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var book models.Book
	id := c.Param("id")
	bookHandler.DB.First(&book, id)
	c.JSON(http.StatusOK, book)
}

func (handler  *Handler) DelBook(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var book models.Book
	id := c.Param("id")
	bookHandler.DB.Delete(&book, id)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}

func (handler  *Handler) UpdateBook(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var book models.Book
	id := c.Param("id")
	bookHandler.DB.First(&book, id)
	json.NewDecoder(c.Request.Body).Decode(&book)
	bookHandler.DB.Save(&book)
	c.JSON(http.StatusOK, book)
}
