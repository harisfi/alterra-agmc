package migration

import (
	"log"

	"github.com/harisfi/alterra-agmc/day6/internal/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.User{},
		&model.Book{},
	)

	if err != nil {
		panic("failed to migrate tables")
	} else {
		log.Println("tables successfully migrated")
	}
}
