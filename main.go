// Author : Dilip Chauhan
package main

import (
	"github.com/gin-gonic/gin"
)

//the main executatble file
//gin router 
var router *gin.Engine

func main() {
	router = gin.Default()

	InitializeRoutes()

	router.Run(":8081")
}
