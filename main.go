package main

import (
	"net/http"

	"github.com/testapi/handlers"
	"github.com/testapi/models"

	"github.com/gin-gonic/gin"

	docs "github.com/testapi/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var handler handlers.Handler

func main() {
	handler.Connect("localhost", "postgres", "root", "forapi", "5433")

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
