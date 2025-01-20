package main

import (
	"fmt"
	"gin-ecommerce/handlers"
	"gin-ecommerce/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Use(cors.Default())

	r.POST("/Signup", handlers.Signup)

	r.GET("/getuser", handlers.Getuser)

	r.POST("/login", handlers.LoginUser)

	r.POST("/logout", middleware.JWTAuthMiddleware(), handlers.LogoutUser)

	fmt.Println("Running on http//:localhost:8855")
	r.Run(":8855")
}
