package database

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gin-ecommerce/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var databaseInstance *gorm.DB

func InIt() *gorm.DB {

	var err error
	databaseInstance, err = connectionDatabase()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	err = performMigration()
	if err != nil {
		log.Fatalf("Could not auto migrate: %v", err)
	}

	return databaseInstance
}

func connectionDatabase() (*gorm.DB, error) {

	dbUsername := getEnv("DB_USERNAME", "root")
	dbPassword := getEnv("DB_PASSWORD", "root")
	dbName := getEnv("DB_NAME", "ecommerce")
	dbHost := getEnv("DB_HOST", "breezy-gifts-push.loca.lt") // Updated to remove https://
	dbPort := getEnv("DB_PORT", "3306")

	// Build connection string without https:// and using tcp protocol
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbName)

	// Connect to the MySQL database
	databaseConnection, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Enable logging for debugging
	})

	if err != nil {
		return nil, err
	}

	sqlDatabase, err := databaseConnection.DB()
	if err != nil {
		return nil, err
	}

	// Connection pool settings from environment variables
	maxIdleConns, _ := strconv.Atoi(getEnv("DB_MAX_IDLE_CONNS", "10"))
	maxOpenConns, _ := strconv.Atoi(getEnv("DB_MAX_OPEN_CONNS", "25"))
	connMaxLifetime, _ := strconv.Atoi(getEnv("DB_CONN_MAX_LIFETIME", "3600")) // in seconds

	sqlDatabase.SetMaxIdleConns(maxIdleConns)
	sqlDatabase.SetMaxOpenConns(maxOpenConns)
	sqlDatabase.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Second)

	return databaseConnection, nil
}

func performMigration() error {
	err := databaseInstance.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}
	return nil
}

func getEnv(key string, defaultVaule string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultVaule
	}
	return value
}
