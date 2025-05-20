package model
import (
	"time"
)
type UserNews struct {
    ID        uint      `gorm:"primaryKey"`
    Name      string    `gorm:"not null"`
    Username  string    `gorm:"unique"`
    Email     *string   `gorm:"type:varchar(100)"`
    Password  string    `gorm:"not null"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
    UpdatedAT time.Time `gorm:"autoUpdateTime"`
}
