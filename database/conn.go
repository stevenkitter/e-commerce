package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
)

const (
	Default = "115.159.222.199:3306"
)

type Database struct {
	Path    string
	Mariadb *gorm.DB
}

func NewDatabase(path string, mariadb *gorm.DB) *Database {
	if path == "" {
		path = Default
	}
	db := &Database{
		Path:    path,
		Mariadb: mariadb,
	}
	go db.Migrate()
	return db
}

func connectDB(user, password, database string) (*gorm.DB, error) {
	sqlUrl := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=True&loc=Local", user, password, os.Getenv("MARIADB_URL"), database)
	return gorm.Open("mysql", sqlUrl)
}

func SkillsConn(password string) (*gorm.DB, error) {
	return connectDB("skills", password, "skills")
}
