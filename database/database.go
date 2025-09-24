package database

import (
    "fmt"
    "log"
    "os"

    "github.com/ayushwar/ecommerce/models"
    "github.com/joho/godotenv"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
    // Load .env file
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using system environment")
    }

    user := os.Getenv("DB_USER")
    pass := os.Getenv("DB_PASS")
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        user, pass, host, port, dbName)

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    DB = db

    log.Println("Database connected successfully")

    // Call migrate function after connection
    if err := MigrateDB(); err != nil {
        log.Fatal("Failed to migrate database:", err)
    }
}

func MigrateDB() error {
    return DB.AutoMigrate(
        &models.User{},
        &models.Product{},
        &models.Order{},
        &models.OrderItem{},
        &models.CartItem{},
    )
}
