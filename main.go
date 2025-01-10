package main

import (
	"fmt"
	"gin-ecommerce/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Use(cors.Default())

	r.POST("/Signup", handlers.Signup)

	r.GET("/getuser", handlers.Getuser)

	r.POST("/login", handlers.LoginUser)

	fmt.Println("Running on http//:localhost:8855")
	r.Run(":8855")
}
