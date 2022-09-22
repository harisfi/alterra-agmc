package migration

import (
	"log"

	"github.com/harisfi/alterra-agmc/day7/database"
	"github.com/harisfi/alterra-agmc/day7/internal/model"
)

func Migrate() {
	conn := database.GetConnection()

	err := conn.AutoMigrate(
		&model.User{},
		&model.Book{},
	)

	if err != nil {
		panic("failed to migrate tables")
	} else {
		log.Println("tables successfully migrated")
	}
}
