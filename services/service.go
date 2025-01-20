package services

import (
	"errors"
	"fmt"
	"gin-ecommerce/database"
	"gin-ecommerce/models"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
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

// func LoginUser(credentials models.LoginCredentials) (models.User, error) {
// 	databaseInstance := database.InIt()

// 	var user models.User
// 	result := databaseInstance.Where("username = ?", credentials.Username).First(&user)

// 	if result.Error != nil {
// 		log.Println("Login failed: invalid username or password")
// 		return models.User{}, errors.New("invalid username or password")
// 	}

// 	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
// 	if err != nil {
// 		log.Println("Login failed: invalid password")
// 		return models.User{}, errors.New(" username or password")
// 	}

// 	return user, nil
// }

func LoginUser(credentials models.LoginCredentials) (models.User, string, error) {
	databaseInstance := database.InIt()

	var user models.User
	result := databaseInstance.Where("username = ?", credentials.Username).First(&user)

	if result.Error != nil {
		log.Println("Login failed: invalid username or password")
		return models.User{}, "", errors.New("invalid username or password")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		log.Println("Login failed: invalid password")
		return models.User{}, "", errors.New("invalid username or password")
	}

	// Check if user already has an active session
	if user.IsLoggedIn {
		log.Println("User is already logged in on another device.")
		return models.User{}, "", errors.New("user already logged in on another device")
	}

	// Generate JWT token for the session
	token, err := generateJWT(user)
	if err != nil {
		log.Println("Error generating token:", err)
		return models.User{}, "", err
	}

	// Mark user as logged in
	user.IsLoggedIn = true
	databaseInstance.Save(&user)

	return user, token, nil
}

func generateJWT(user models.User) (string, error) {
	// Create a JWT token
	claims := jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Use your secret key here
	secretKey := []byte("your_secret_key")

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func LogoutUser(userId uint) error {
	databaseInstance := database.InIt()

	var user models.User
	result := databaseInstance.First(&user, userId)

	if result.Error != nil {
		log.Println("Logout failed: user not found")
		return result.Error
	}

	// Invalidate the session by marking the user as logged out
	user.IsLoggedIn = false
	databaseInstance.Save(&user)

	log.Println("User logged out successfully")
	return nil
}
