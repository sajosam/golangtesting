package usermngmnt

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/testapi/dbmngmnt"
)


func HealthCheck(c *gin.Context) {
	c.Status(http.StatusOK)
	c.String(http.StatusOK, "Super Secret Area")
}

func (usrHandler *UsrHandler) GetUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var users []dbmngmnt.User
	usrHandler.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

func (usrHandler *UsrHandler) AddUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var user dbmngmnt.User
	json.NewDecoder(c.Request.Body).Decode(&user)
	usrHandler.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}

func (usrHandler *UsrHandler) GetUserInd(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var user dbmngmnt.User
	id := c.Param("id")
	usrHandler.DB.First(&user, id)
	c.JSON(http.StatusOK, user)
}

func (usrHandler *UsrHandler) DelUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var user dbmngmnt.User
	id := c.Param("id")
	usrHandler.DB.Delete(&user, id)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

func (usrHandler *UsrHandler) UpdateUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var user dbmngmnt.User
	id := c.Param("id")
	usrHandler.DB.First(&user, id)
	json.NewDecoder(c.Request.Body).Decode(&user)
	usrHandler.DB.Save(&user)
	c.JSON(http.StatusOK, user)
}
