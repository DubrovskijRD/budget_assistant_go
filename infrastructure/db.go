package infrastructure

import (
	"fmt"

	"github.com/DubrovskijRD/budget_assistant_go/application"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database interface {
	GetDb() *gorm.DB
	Close()
}

type postgresDatabase struct {
	Db *gorm.DB
}

func NewPostgresDatabase(cfg *application.Config) Database {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBName,
		cfg.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return &postgresDatabase{Db: db}
}

func (p *postgresDatabase) GetDb() *gorm.DB {
	return p.Db
}

func (p *postgresDatabase) Close() {
	// todo for graseful shutdown
}
