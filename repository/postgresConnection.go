package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func PostgresConnection() (*PostgresRepository, error) {
	db, err := gorm.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", "127.0.0.1", "5432", "postgres", "sinkar", "Go4thsas", "disable"))

	if err != nil {
		return nil, err
	}
	return &PostgresRepository{
		db,
	}, nil
}
