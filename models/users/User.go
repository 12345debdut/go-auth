package users

import (
	"time"

	"gorm.io/gorm"
)

// User represents a registered user in the system.
type User struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string         `gorm:"size:100;not null;uniqueIndex" json:"username"`
	Email     string         `gorm:"size:255;not null;uniqueIndex" json:"email"`
	Password  string         `gorm:"size:255;not null" json:"-"`
	FirstName string         `gorm:"size:100" json:"first_name,omitempty"`
	LastName  string         `gorm:"size:100" json:"last_name,omitempty"`
	IsActive  bool           `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
