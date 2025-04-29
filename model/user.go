package model 

type Users struct {
    ID       uint   `gorm:"primaryKey"`
    Username string `gorm:"unique"`
    Email    string
}