package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
	"fmt"
	"database/sql"
)

type DbConfig struct {
	Host string
	Port string
	User string
	Password string
	Name string
}

func CreateConnection(cfg DbConfig) (*sql.DB, error) {
	log.Debugf("Connecting to MySQL on '%s' to database '%s'", cfg.Host, cfg.Name)
	return sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name))
}
