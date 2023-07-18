package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)


func TestAddUser(t *testing.T) {
    router := gin.Default()
    router.POST("/adduser", mainhadler.AddUser)

    req, _ := http.NewRequest("POST", "/adduser", strings.NewReader(`{"ID": 15, "user_name": "Alice", "email": "alice@example.com"}`))
    req.Header.Set("Content-Type", "application/json")
    recorder := httptest.NewRecorder()

    router.ServeHTTP(recorder, req)

    fmt.Println("AddUser Test Result:", recorder.Body.String())
    assert.Equal(t, http.StatusCreated, recorder.Code)
}

func TestGetUser(t *testing.T) {
    router := gin.Default()
    router.GET("/user", mainhadler.GetUser)

    req, _ := http.NewRequest("GET", "/user", nil)
    recorder := httptest.NewRecorder()

    router.ServeHTTP(recorder, req)

    fmt.Println("GetUser Test Result:", recorder.Body.String())
    assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestGetUserInd(t *testing.T) {
    router := gin.Default()
    router.GET("/user/:id", mainhadler.GetUserInd)

    req, _ := http.NewRequest("GET", "/user/15", nil)
    recorder := httptest.NewRecorder()

    router.ServeHTTP(recorder, req)

    fmt.Println("GetUserInd Test Result:", recorder.Body.String())
    assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestUpdateUser(t *testing.T) {
    router := gin.Default()
    router.PUT("/updateUser/:id", mainhadler.UpdateUser)

    req, _ := http.NewRequest("PUT", "/updateUser/15", strings.NewReader(`{"user_name": "Alice-updated", "email": "alice.smith@example.com"}`))
    req.Header.Set("Content-Type", "application/json")
    recorder := httptest.NewRecorder()

    router.ServeHTTP(recorder, req)

    fmt.Println("UpdateUser Test Result:", recorder.Body.String())
    assert.Equal(t, http.StatusOK, recorder.Code)
}


func TestDelUser(t *testing.T) {
    router := gin.Default()
    router.DELETE("/delUser/:id", mainhadler.DelUser)

    req, _ := http.NewRequest("DELETE", "/delUser/15", nil)
    recorder := httptest.NewRecorder()

    router.ServeHTTP(recorder, req)

    fmt.Println("DelUser Test Result:", recorder.Body.String())
    assert.Equal(t, http.StatusOK, recorder.Code)
}

