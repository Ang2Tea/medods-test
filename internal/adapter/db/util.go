package db

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UUIDModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func Migration(db *gorm.DB) error {
	db.DisableForeignKeyConstraintWhenMigrating = true

	err := db.AutoMigrate(new(User))
	if err != nil {
		return err
	}

	db.DisableForeignKeyConstraintWhenMigrating = false

	return db.AutoMigrate(new(User))
}
