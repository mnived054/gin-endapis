package handlers

import (
	"gin-ecommerce/models"
	"gin-ecommerce/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	var user models.User
	if c.BindJSON(&user) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
	}

	createduser, err := services.CreateUser(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"user created": createduser})
	}
}

func Getuser(c *gin.Context) {
	users, err := services.GetAllUser()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func LoginUser(c *gin.Context) {
	var credentials models.LoginCredentials

	if err := c.ShouldBindJSON(&credentials); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	log.Println("Received credentials:", credentials)

	user, err := services.LoginUser(credentials)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Successfully login": user})

}
