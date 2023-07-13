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
	Id int `json:"ID"`
	User_name string `json:"user_name"`
	Email string `json:"email"`
}

type UsrHandler struct {
	DB *gorm.DB
}

func (ordhandler *UsrHandler) Connection(host,user,password,dbname,port string) {
	var err error

	dsn:="host="+host+" user="+user+" password="+password+" dbname="+dbname+" port="+port+" sslmode=disable"
	ordhandler.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")

	// ordhandler.DB.AutoMigrate(User{})
	ordhandler.DB.AutoMigrate(&User{})

}

// func HealthCheck(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintf(w, "Super Secret Area")
// }

func HealthCheck(c *gin.Context) {
	c.Status(http.StatusOK)
	c.String(http.StatusOK, "Super Secret Area")
}



// func (usrhandler *UsrHandler) GetUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var User []User
// 	usrhandler.DB.Find(&User)
// 	json.NewEncoder(w).Encode(&User)
// }

func (usrhandler *UsrHandler) GetUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var users []User
	usrhandler.DB.Find(&users)
	c.JSON(200, users)
}


// func (usrhandler *UsrHandler) AddUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var User User
// 	json.NewDecoder(r.Body).Decode(&User)
// 	usrhandler.DB.Create(&User)
// 	json.NewEncoder(w).Encode(&User)
// }

func (usrhandler *UsrHandler) AddUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var user User
	json.NewDecoder(c.Request.Body).Decode(&user)
	usrhandler.DB.Create(&user)
	c.JSON(201, user)
}


// func (usrHandler *UsrHandler) GetUserInd(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var user User
// 	params := mux.Vars(r)
// 	usrHandler.DB.First(&user, params["id"])
// 	json.NewEncoder(w).Encode(&user)

// }

func (usrHandler *UsrHandler) GetUserInd(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var user User
	id := c.Param("id")
	usrHandler.DB.First(&user, id)
	c.JSON(200, user)
}


// func (usrhandler *UsrHandler) DelUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var User User
// 	params := mux.Vars(r)
// 	usrhandler.DB.Delete(&User, params["id"])
// 	json.NewEncoder(w).Encode(&User)
	
// }

func (usrhandler *UsrHandler) DelUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var user User
	id := c.Param("id")
	usrhandler.DB.Delete(&user, id)
	c.JSON(200, gin.H{"message": "User deleted"})
}


// func (usrhandler *UsrHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var User User
// 	params := mux.Vars(r)
// 	usrhandler.DB.First(&User, params["id"])
// 	json.NewDecoder(r.Body).Decode(&User)
// 	usrhandler.DB.Save(&User)
// 	json.NewEncoder(w).Encode(&User)
	
// }

func (usrhandler *UsrHandler) UpdateUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var user User
	id := c.Param("id")
	usrhandler.DB.First(&user, id)
	json.NewDecoder(c.Request.Body).Decode(&user)
	usrhandler.DB.Save(&user)
	c.JSON(200, user)
}
