package seeder

import (
	"log"

	"github.com/harisfi/alterra-agmc/day10/internal/model"
	"gorm.io/gorm"
)

func bookSeeder(conn *gorm.DB) {
	var books = []model.Book{
		{Title: "Book A", Author: model.User{IDModel: model.IDModel{ID: 1}}, Publisher: "Publisher A"},
		{Title: "Book B", Author: model.User{IDModel: model.IDModel{ID: 1}}, Publisher: "Publisher B"},
		{Title: "Book C", Author: model.User{IDModel: model.IDModel{ID: 2}}, Publisher: "Publisher C"},
	}

	if err := conn.Create(&books).Error; err != nil {
		log.Printf("failed to seed user data\n%v\n", err)
	}

	log.Println("successfully seeding user data")
}
