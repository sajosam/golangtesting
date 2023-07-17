package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/testapi/models"
)



func (handler  *Handler) GetBook(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var book []models.Book
	handler.DB.Find(&book)
	c.JSON(http.StatusOK, book)
}

func (handler  *Handler) AddBook(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var book models.Book
	json.NewDecoder(c.Request.Body).Decode(&book)
	handler.DB.Create(&book)
	c.JSON(http.StatusCreated, book)
}

func (handler  *Handler) GetBookInd(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var book models.Book
	id := c.Param("id")
	handler.DB.First(&book, id)
	c.JSON(http.StatusOK, book)
}

func (handler  *Handler) DelBook(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var book models.Book
	id := c.Param("id")
	handler.DB.Delete(&book, id)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}

func (handler  *Handler) UpdateBook(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var book models.Book
	id := c.Param("id")
	handler.DB.First(&book, id)
	json.NewDecoder(c.Request.Body).Decode(&book)
	handler.DB.Save(&book)
	c.JSON(http.StatusOK, book)
}
