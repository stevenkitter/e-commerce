package database

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Session struct {
	gorm.Model
	UserID string `gorm:"unique_index"`
	Token  string `gorm:"unique"`
	Login  *time.Time
	Logout *time.Time
}
