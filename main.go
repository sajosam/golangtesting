package main

import (
	"net/http"
	"os"

	"github.com/testapi/handlers"
	"github.com/testapi/models"

	"github.com/gin-gonic/gin"

	docs "github.com/testapi/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var handler handlers.Handler

func main() {

	// set environment variables
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5433")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "root")
	os.Setenv("DB_NAME", "forapi")

	os.Setenv("DB_HOST_TEST", "localhost")
	os.Setenv("DB_PORT_TEST", "5433")
	os.Setenv("DB_USER_TEST", "postgres")
	os.Setenv("DB_PASSWORD_TEST", "root")
	os.Setenv("DB_NAME_TEST", "testapi")


	handler.Connect(os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	handler.DB.AutoMigrate(&models.Book{})
	handler.DB.AutoMigrate(&models.User{})

	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"

	router.GET("/userhealth", handlers.HealthCheck)
	router.GET("/user", handler.GetUser)
	router.POST("/adduser", handler.AddUser)
	router.GET("/user/:id", handler.GetUserInd)
	router.DELETE("/delUser/:id", handler.DelUser)
	router.PUT("/updateUser/:id", handler.UpdateUser)


	router.GET("/book", handler.GetBook)
	router.POST("/addbook", handler.AddBook)
	router.GET("/book/:id", handler.GetBookInd)
	router.DELETE("/delBook/:id", handler.DelBook)
	router.PUT("/updateBook/:id", handler.UpdateBook)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	http.Handle("/", router)
	http.ListenAndServe("0.0.0.0:8000", router)

	dbInstance, _ := handler.DB.DB()
	defer dbInstance.Close()
}
