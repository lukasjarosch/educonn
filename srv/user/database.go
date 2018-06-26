package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
	"fmt"
)

type DbConfig struct {
	Host string
	Port string
	User string
	Password string
	Name string
}

func CreateConnection(cfg DbConfig) (*gorm.DB, error) {
	log.Debugf("Connecting to MySQL on '%s' to database '%s'", cfg.Host, cfg.Name)
	return gorm.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name))
}
