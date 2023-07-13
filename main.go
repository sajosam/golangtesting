package main

import (
	"net/http"

	"example/userapi/usermngmnt"

	"github.com/gin-gonic/gin"

	docs "example/userapi/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)


var usrhandlerobj usermngmnt.UsrHandler
func main() {
	usrhandlerobj.Connection("localhost","postgres","root","forapi","5433")

	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"

	router.GET("/health", usermngmnt.HealthCheck)
	router.GET("/user", usrhandlerobj.GetUser)
	router.POST("/adduser", usrhandlerobj.AddUser)
	router.GET("/user/:id", usrhandlerobj.GetUserInd)
	router.DELETE("/delUser/:id", usrhandlerobj.DelUser)
	router.PUT("/updateUser/:id", usrhandlerobj.UpdateUser)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	http.Handle("/", router)
	http.ListenAndServe("0.0.0.0:8000", router)

	dbinstance,_ := usrhandlerobj.DB.DB()
	defer dbinstance.Close()
}

