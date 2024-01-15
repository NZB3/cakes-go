package repository

import (
	"github.com/nzb3/cakes-go/internal/lib/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type repository struct {
	db  *gorm.DB
	log logger.Logger
}

var instance *repository

func NewRepository(log logger.Logger) *repository {
	if instance != nil {
		return instance
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("POSTGRES_DSN")), &gorm.Config{})
	if err != nil {
		log.Errorf("error creating repository: %w", err)
		panic(err)
	}

	instance = &repository{
		db:  db,
		log: log,
	}

	return instance
}
