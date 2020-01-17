// Author : Dilip Chauhan
package main

import (
	"github.com/gin-gonic/gin"
)

//gin router 
//good download manager
//good
var router *gin.Engine

func main() {
	router = gin.Default()

	InitializeRoutes()

	router.Run(":8081")
}
