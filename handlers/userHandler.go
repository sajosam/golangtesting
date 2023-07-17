package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/testapi/models"
)


func HealthCheck(c *gin.Context) {
	c.Status(http.StatusOK)
	c.String(http.StatusOK, "Super Secret Area")
}

func (handler Handler) GetUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var users []models.User
	handler.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

func (handler Handler) AddUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var user models.User
	json.NewDecoder(c.Request.Body).Decode(&user)
	handler.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}

func (handler Handler) GetUserInd(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var user models.User
	id := c.Param("id")
	handler.DB.First(&user, id)
	c.JSON(http.StatusOK, user)
}

func (handler Handler) DelUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var user models.User
	id := c.Param("id")
	handler.DB.Delete(&user, id)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

func (handler Handler) UpdateUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var user models.User
	id := c.Param("id")
	handler.DB.First(&user, id)
	json.NewDecoder(c.Request.Body).Decode(&user)
	handler.DB.Save(&user)
	c.JSON(http.StatusOK, user)
}
