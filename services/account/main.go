package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/stevenkitter/skills/database"
	"github.com/stevenkitter/skills/services/account/server"
	"os"
)

var (
	dbPath                = os.Getenv("MARIADB_URL")
	skillsMariadbPassword = os.Getenv("SKILLS_MARIADB_PWD")
	accountLogger         *log.Entry
)

const (
	port = "51000"
)

func init() {
	accountLogger = log.WithFields(log.Fields{
		"server": "account",
		"port":   port,
	})
}

func main() {
	gorm, err := database.SkillsConn(skillsMariadbPassword)
	if err != nil {
		accountLogger.Panicf("database.SkillsConn() err:%v", err)
	}
	defer func() {
		if err := gorm.Close(); err != nil {
			accountLogger.Panicf("gorm.Close() err:%v", err)
		}
	}()
	db := database.NewDatabase(dbPath, gorm)
	engine := server.Engine{
		Name: "micro.server.account",
		Port: port,
		DB:   db,
	}
	if err := engine.Run(); err != nil {
		accountLogger.Panicf("engine.Run() err:%v", err)
	}
}
