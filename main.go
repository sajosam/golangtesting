package main

import (
	"net/http"

	"github.com/testapi/usermngmnt"

	"github.com/gin-gonic/gin"

	docs "github.com/testapi/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var usrHandlerObj usermngmnt.UsrHandler

func main() {
	usrHandlerObj.Connection("localhost", "postgres", "root", "forapi", "5433")

	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"

	router.GET("/userhealth", usermngmnt.HealthCheck)
	router.GET("/user", usrHandlerObj.GetUser)
	router.POST("/adduser", usrHandlerObj.AddUser)
	router.GET("/user/:id", usrHandlerObj.GetUserInd)
	router.DELETE("/delUser/:id", usrHandlerObj.DelUser)
	router.PUT("/updateUser/:id", usrHandlerObj.UpdateUser)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	http.Handle("/", router)
	http.ListenAndServe("0.0.0.0:8000", router)

	dbInstance, _ := usrHandlerObj.DB.DB()
	defer dbInstance.Close()
}
