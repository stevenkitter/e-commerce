package server

import "github.com/jinzhu/gorm"

type Engine struct {
	Name string
	Port string
	DB   *gorm.DB
}

type Server interface {
	Run() error
}
