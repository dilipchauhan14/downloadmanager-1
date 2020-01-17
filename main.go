// Author : Dilip Chauhan
package main

import (
	"github.com/gin-gonic/gin"
)
//A download manager is a software tool that manages the downloading of files from the Internet, which may be built: 
//into a Web browser, or as a, usually more sophisticated, stand-alone manager.
//gin router 
var router *gin.Engine

func main() {
	router = gin.Default()

	InitializeRoutes()

	router.Run(":8081")
}
