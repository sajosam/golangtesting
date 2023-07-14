package bookmngmnt

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

var bookHandler BookHandler

func TestMain(m *testing.M) {
	dsn := "host=localhost user=postgres password=root dbname=forapi port=5433 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
	bookHandler.DB = db
    bookHandler.BookConnection("localhost", "postgres", "root", "forapi", "5433")

    code := m.Run()

    dbInstance, _ := bookHandler.DB.DB()
    dbInstance.Close()

    os.Exit(code)
}

func TestConnection(t *testing.T) {
	assert.NotNil(t, bookHandler.DB)
}

func TestGetBook(t *testing.T) {
    router := gin.Default()
    router.GET("/book", bookHandler.GetBook)

    req, _ := http.NewRequest("GET", "/book", nil)
    recorder := httptest.NewRecorder()

    router.ServeHTTP(recorder, req)

    fmt.Println("GetBook Test Result:", recorder.Body.String())
    assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestAddBook(t *testing.T) {
    router := gin.Default()
    router.POST("/addbook", bookHandler.AddBook)

    req, _ := http.NewRequest("POST", "/addbook", strings.NewReader(`{"ID": 10, "book_name": "Alice", "price": 100}`))
    req.Header.Set("Content-Type", "application/json")
    recorder := httptest.NewRecorder()

    router.ServeHTTP(recorder, req)

    fmt.Println("AddBook Test Result:", recorder.Body.String())
    assert.Equal(t, http.StatusCreated, recorder.Code)
}

func TestGetBookInd(t *testing.T) {
    router := gin.Default()
    router.GET("/book/:id", bookHandler.GetBookInd)

    req, _ := http.NewRequest("GET", "/book/10", nil)
    recorder := httptest.NewRecorder()

    router.ServeHTTP(recorder, req)

    fmt.Println("GetBookInd Test Result:", recorder.Body.String())
    assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestUpdateBook(t *testing.T) {
    router := gin.Default()
    router.PUT("/updateBook/:id", bookHandler.UpdateBook)

    req, _ := http.NewRequest("PUT", "/updateBook/10", strings.NewReader(`{"book_name": "Alice", "price": 1000}`))
    req.Header.Set("Content-Type", "application/json")
    recorder := httptest.NewRecorder()

    router.ServeHTTP(recorder, req)

    fmt.Println("UpdateBook Test Result:", recorder.Body.String())
    assert.Equal(t, http.StatusOK, recorder.Code)
}


func TestDelBook(t *testing.T) {
    router := gin.Default()
    router.DELETE("/delBook/:id", bookHandler.DelBook)

    req, _ := http.NewRequest("DELETE", "/delBook/10", nil)
    recorder := httptest.NewRecorder()

    router.ServeHTTP(recorder, req)

    fmt.Println("DelUser Test Result:", recorder.Body.String())
    assert.Equal(t, http.StatusOK, recorder.Code)
}

