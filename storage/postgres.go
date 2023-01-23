package storage

import (
	"fmt"

	"gorm.io/driver/postgres"
)

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBname   string
}

func NewConnection(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("Host=%s Port=%s Password=%s User=%s DBname=%s", config.Host, config.Port, config.Password, config.Password, config.User, config.DBname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	return db, nil
}
