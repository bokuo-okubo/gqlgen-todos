package infra

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/bokuo-okubo/gqlgen-todos/config"
	"github.com/bokuo-okubo/gqlgen-todos/entity"
)

var DbConnection *sql.DB

func ConnectDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(config.Config.DbName), &gorm.Config{})
	fmt.Printf("%+v\n", config.Config)
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&entity.User{}, &entity.Todo{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
