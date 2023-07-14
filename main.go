package main

import (
	"net/http"

	"github.com/testapi/bookmngmnt"
	"github.com/testapi/usermngmnt"

	"github.com/gin-gonic/gin"

	docs "github.com/testapi/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var usrHandlerObj usermngmnt.UsrHandler
var bookHandlerObj bookmngmnt.BookHandler

func main() {
	usrHandlerObj.UserConnection("localhost", "postgres", "root", "forapi", "5433")
	bookHandlerObj.BookConnection("localhost", "postgres", "root", "forapi", "5433")


	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"

	router.GET("/userhealth", usermngmnt.HealthCheck)
	router.GET("/user", usrHandlerObj.GetUser)
	router.POST("/adduser", usrHandlerObj.AddUser)
	router.GET("/user/:id", usrHandlerObj.GetUserInd)
	router.DELETE("/delUser/:id", usrHandlerObj.DelUser)
	router.PUT("/updateUser/:id", usrHandlerObj.UpdateUser)


	router.GET("/bookhealth", bookmngmnt.HealthCheck)
	router.GET("/book", bookHandlerObj.GetBook)
	router.POST("/addbook", bookHandlerObj.AddBook)
	router.GET("/book/:id", bookHandlerObj.GetBookInd)
	router.DELETE("/delBook/:id", bookHandlerObj.DelBook)
	router.PUT("/updateBook/:id", bookHandlerObj.UpdateBook)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	http.Handle("/", router)
	http.ListenAndServe("0.0.0.0:8000", router)

	dbInstance, _ := usrHandlerObj.DB.DB()
	defer dbInstance.Close()
}
