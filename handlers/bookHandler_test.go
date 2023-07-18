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

func TestAddBook(t *testing.T) {
    router := gin.Default()
    router.POST("/addbook", mainhadler.AddBook)

    req, _ := http.NewRequest("POST", "/addbook", strings.NewReader(`{"ID": 1, "book_name": "newbook", "price": 100}`))
    req.Header.Set("Content-Type", "application/json")
    recorder := httptest.NewRecorder()

    router.ServeHTTP(recorder, req)

    fmt.Println("AddBook Test Result:", recorder.Body.String())
    assert.Equal(t, http.StatusCreated, recorder.Code)
}

func TestGetBook(t *testing.T) {
    router := gin.Default()
    router.GET("/book", mainhadler.GetBook)

    req, _ := http.NewRequest("GET", "/book", nil)
    recorder := httptest.NewRecorder()
    router.ServeHTTP(recorder, req)

    fmt.Println("GetBook Test Result:", recorder.Body.String())
    assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestGetBookInd(t *testing.T) {
    router := gin.Default()
    router.GET("/book/:id", mainhadler.GetBookInd)

    req, _ := http.NewRequest("GET", "/book/1", nil)
    recorder := httptest.NewRecorder()

    router.ServeHTTP(recorder, req)

    fmt.Println("GetBookInd Test Result:", recorder.Body.String())
    assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestUpdateBook(t *testing.T) {
    router := gin.Default()
    router.PUT("/updateBook/:id", mainhadler.UpdateBook)

    req, _ := http.NewRequest("PUT", "/updateBook/1", strings.NewReader(`{"book_name": "Alice-updated", "price": 1000}`))
    req.Header.Set("Content-Type", "application/json")
    recorder := httptest.NewRecorder()

    router.ServeHTTP(recorder, req)

    fmt.Println("UpdateBook Test Result:", recorder.Body.String())
    assert.Equal(t, http.StatusOK, recorder.Code)
}


func TestDelBook(t *testing.T) {
    router := gin.Default()
    router.DELETE("/delBook/:id", mainhadler.DelBook)

    req, _ := http.NewRequest("DELETE", "/delBook/1", nil)
    recorder := httptest.NewRecorder()

    router.ServeHTTP(recorder, req)

    fmt.Println("DelUser Test Result:", recorder.Body.String())
    assert.Equal(t, http.StatusOK, recorder.Code)
}

