package db

import (
	"fmt"
	"github.com/bperezgo/go-project/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetConnection() *gorm.DB {
	if DB != nil {
		return DB
	}
	env := config.GetConfiguration()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", env.PgHostname, env.PgUser, env.PgPassword, env.PgDB, env.PgPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
	return DB
}
