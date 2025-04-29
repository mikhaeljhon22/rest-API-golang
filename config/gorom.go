package config 
import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "os"
    "github.com/joho/godotenv"
    "log"
    "fmt"
)
func ConnectToPostgreSQL() (*gorm.DB, error) {
    err := godotenv.Load()
    if err != nil{
        log.Fatal("error load env variabel")
    }

    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")
    
    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host,port,user,password,dbname)
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
       panic("failed to connect the database")
    }
    return db, nil
}
