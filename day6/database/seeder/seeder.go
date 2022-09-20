package seeder

import "github.com/harisfi/alterra-agmc/day6/database"

func Seed() {
	conn := database.GetConnection()

	userSeeder(conn)
	bookSeeder(conn)
}
