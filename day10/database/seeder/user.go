package seeder

import (
	"log"

	"github.com/harisfi/alterra-agmc/day10/internal/model"
	"gorm.io/gorm"
)

func userSeeder(conn *gorm.DB) {
	var users = []model.User{
		{Name: "User A", Email: "a@mail.com", Password: "12345678"},
		{Name: "User B", Email: "b@mail.com", Password: "12345678"},
		{Name: "User C", Email: "c@mail.com", Password: "12345678"},
	}

	if err := conn.Create(&users).Error; err != nil {
		log.Printf("failed to seed user data\n%v\n", err)
	}

	log.Println("successfully seeding user data")
}
