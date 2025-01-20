package database

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gin-ecommerce/models"

	"gorm.io/driver/postgres"
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

	dbUsername := getEnv("DB_USERNAME", "nivedinstance_user")               // Default username for PostgreSQL is usually 'postgres'
	dbPassword := getEnv("DB_PASSWORD", "dOB8aDlUAksWthWhnJjyEix0cvSeJS3a") // Default password (change it for your production app)
	dbName := getEnv("DB_NAME", "nivedinstance")                            // The name of your PostgreSQL database
	dbHost := getEnv("DB_HOST", "dpg-cu0jgjaj1k6c73c0ptsg-a")               // The Render-provided PostgreSQL host URL
	dbPort := getEnv("DB_PORT", "5432")                                     // Default PostgreSQL port is 5432

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUsername, dbPassword, dbName)

	databaseConnection, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	sqlDatabase, err := databaseConnection.DB()
	if err != nil {
		return nil, err
	}

	maxIdleConns, _ := strconv.Atoi(getEnv("DB_MAX_IDLE_CONNS", "10"))
	maxOpenConns, _ := strconv.Atoi(getEnv("DB_MAX_OPEN_CONNS", "25"))
	connMaxLifetime, _ := strconv.Atoi(getEnv("DB_CONN_MAX_LIFETIME", "3600"))

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

//My-Sql db
// package database

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// 	"time"

// 	"gin-ecommerce/models"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// )

// var databaseInstance *gorm.DB

// func InIt() *gorm.DB {

// 	var err error
// 	databaseInstance, err = connectionDatabase()
// 	if err != nil {
// 		log.Fatalf("Could not connect to the database: %v", err)
// 	}

// 	err = performMigration()
// 	if err != nil {
// 		log.Fatalf("Could not auto migrate: %v", err)
// 	}

// 	return databaseInstance

// }

// func connectionDatabase() (*gorm.DB, error) {

// 	dbUsername := getEnv("DB_USERNAME", "root")
// 	dbPassword := getEnv("DB_PASSWORD", "root")
// 	dbName := getEnv("DB_NAME", "ecommerce")
// 	dbHost := getEnv("DB_HOST", "localhost")
// 	dbPort := getEnv("DB_PORT", "3306")

// 	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbName)

// 	databaseConnection, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
// 		Logger: logger.Default.LogMode(logger.Info), // Enable logging for debugging
// 	})

// 	if err != nil {
// 		return nil, err
// 	}

// 	sqlDatabase, err := databaseConnection.DB()
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Connection pool settings from environment variables
// 	maxIdleConns, _ := strconv.Atoi(getEnv("DB_MAX_IDLE_CONNS", "10"))
// 	maxOpenConns, _ := strconv.Atoi(getEnv("DB_MAX_OPEN_CONNS", "25"))
// 	connMaxLifetime, _ := strconv.Atoi(getEnv("DB_CONN_MAX_LIFETIME", "3600")) // in seconds

// 	sqlDatabase.SetMaxIdleConns(maxIdleConns)
// 	sqlDatabase.SetMaxOpenConns(maxOpenConns)
// 	sqlDatabase.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Second)

// 	return databaseConnection, nil
// }

// func performMigration() error {
// 	err := databaseInstance.AutoMigrate(&models.User{})
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func getEnv(key string, defaultVaule string) string {
// 	value := os.Getenv(key)
// 	if value == "" {
// 		return defaultVaule
// 	}
// 	return value
// }
