package main

import (
	"gin-ecommerce/models"
	repo "gin-ecommerce/repo"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.POST("/Signup", func(c *gin.Context) {
		var user models.User
		if c.BindJSON(&user) != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		}

		createduser, err := repo.CreateUser(user)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
		} else {
			c.JSON(http.StatusOK, gin.H{"user created": createduser})
		}
	})

	r.GET("/getuser", func(c *gin.Context) {
		users, err := repo.GetAllUser()
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"users": users})
	})

	r.POST("/login", func(c *gin.Context) {
		var credentials models.LoginCredentials

		if err := c.ShouldBindJSON(&credentials); err != nil {
			log.Println("Error binding JSON:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
			return
		}

		log.Println("Received credentials:", credentials)

		user, err := repo.LoginUser(credentials)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"errors": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"Successfully login": user})

	})

	r.Run(":8855")
}
