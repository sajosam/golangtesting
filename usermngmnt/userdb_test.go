package usermngmnt

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var usrHandler UsrHandler

func TestMain(m *testing.M) {
	dsn := "host=localhost user=postgres password=root dbname=forapi port=5433 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
	usrHandler.DB = db
    usrHandler.Connection("localhost", "postgres", "root", "forapi", "5433")

    code := m.Run()

    dbInstance, _ := usrHandler.DB.DB()
    dbInstance.Close()

    os.Exit(code)
}

func TestConnection(t *testing.T) {
	assert.NotNil(t, usrHandler.DB)
}

func TestGetUser(t *testing.T) {
    router := gin.Default()
    router.GET("/user", usrHandler.GetUser)

    req, _ := http.NewRequest("GET", "/user", nil)
    recorder := httptest.NewRecorder()

    router.ServeHTTP(recorder, req)

    fmt.Println("GetUser Test Result:", recorder.Body.String())
    assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestAddUser(t *testing.T) {
    router := gin.Default()
    router.POST("/adduser", usrHandler.AddUser)

    req, _ := http.NewRequest("POST", "/adduser", strings.NewReader(`{"ID": 12, "user_name": "Alice", "email": "alice@example.com"}`))
    req.Header.Set("Content-Type", "application/json")
    recorder := httptest.NewRecorder()

    router.ServeHTTP(recorder, req)

    fmt.Println("AddUser Test Result:", recorder.Body.String())
    assert.Equal(t, http.StatusCreated, recorder.Code)
}

func TestGetUserInd(t *testing.T) {
    router := gin.Default()
    router.GET("/user/:id", usrHandler.GetUserInd)

    req, _ := http.NewRequest("GET", "/user/12", nil)
    recorder := httptest.NewRecorder()

    router.ServeHTTP(recorder, req)

    fmt.Println("GetUserInd Test Result:", recorder.Body.String())
    assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestUpdateUser(t *testing.T) {
    router := gin.Default()
    router.PUT("/updateUser/:id", usrHandler.UpdateUser)

    req, _ := http.NewRequest("PUT", "/updateUser/12", strings.NewReader(`{"user_name": "Alice", "email": "alice.smith@example.com"}`))
    req.Header.Set("Content-Type", "application/json")
    recorder := httptest.NewRecorder()

    router.ServeHTTP(recorder, req)

    fmt.Println("UpdateUser Test Result:", recorder.Body.String())
    assert.Equal(t, http.StatusOK, recorder.Code)
}


func TestDelUser(t *testing.T) {
    router := gin.Default()
    router.DELETE("/delUser/:id", usrHandler.DelUser)

    req, _ := http.NewRequest("DELETE", "/delUser/12", nil)
    recorder := httptest.NewRecorder()

    router.ServeHTTP(recorder, req)

    fmt.Println("DelUser Test Result:", recorder.Body.String())
    assert.Equal(t, http.StatusOK, recorder.Code)
}

