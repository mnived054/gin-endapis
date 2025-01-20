package services

import (
	"errors"
	"fmt"
	"gin-ecommerce/database"
	"gin-ecommerce/models"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user models.User) (models.User, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		return models.User{}, err
	}

	user.Password = string(hashedPassword)

	databaseInstance := database.InIt()

	var existingUser models.User
	result := databaseInstance.Where("username = ?", user.Username).First(&existingUser)

	if result.Error == nil && existingUser.ID != user.ID {
		log.Println("Username already taken,  but with a different ID.")
		return models.User{}, fmt.Errorf("username %s is already taken by a different user", user.Username)
	}

	result = databaseInstance.Where("id = ?", user.ID).First(&existingUser)
	if result.Error == nil {
		log.Println("User with the same ID already exists.")
		return models.User{}, fmt.Errorf("user with the same ID already exists")
	}

	result = databaseInstance.Create(&user)
	if result.Error != nil {
		log.Println("Error during user creation", result.Error)
		return models.User{}, result.Error
	}

	return user, nil
}

func GetAllUser() ([]models.User, error) {

	var existingUser []models.User
	databaseInstance := database.InIt()
	databaseInstance.Find(&existingUser)

	return existingUser, nil
}

func LoginUser(credentials models.LoginCredentials) (models.User, error) {
	databaseInstance := database.InIt()

	var user models.User
	result := databaseInstance.Where("username = ?", credentials.Username).First(&user)

	if result.Error != nil {
		log.Println("Login failed: invalid username or password")
		return models.User{}, errors.New("invalid username or password")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		log.Println("Login failed: invalid password")
		return models.User{}, errors.New(" username or password")
	}

	return user, nil
}
