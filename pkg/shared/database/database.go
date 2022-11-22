package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

type Database struct {
	*gorm.DB
}

func New(config ConfigDatabase) *Database {
	fmt.Println("Try NewDatabase ...")

	db, err := gorm.Open("postgres",
		fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%v sslmode=disable search_path=%s",
			config.Username,
			config.Password,
			config.Name,
			config.Host,
			config.Port,
			config.Schema))

	if err != nil {
		fmt.Println("failed to connect database")
		panic(err)
	}

	db.SingularTable(false)

	err = db.DB().Ping()
	if err != nil {
		fmt.Println("failed to connect database")
		panic(err)
	}

	db.DB().SetMaxIdleConns(config.MinIdleConnections)
	db.DB().SetMaxOpenConns(config.MaxOpenConnections)
	db.DB().SetConnMaxLifetime(config.ConnMaxLifetime)

	db.LogMode(config.DebugMode)

	s := &Database{
		db,
	}

	if err = s.MigrateUP("./migration"); err != nil {
		fmt.Println("failed to migrate database")
		panic(err)
	}

	return s
}

func (s *Database) MigrateUP(path string) error {
	sqlDb := s.DB.DB()
	if err := goose.Up(sqlDb, path); err != nil {
		return fmt.Errorf("migration: %w", err)
	}

	return nil
}
