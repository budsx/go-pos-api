package config

import (
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection() *gorm.DB {
	db, err := gorm.Open(postgres.Open("postgres://postgres:21799@localhost:5432/go-pos-api?sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
