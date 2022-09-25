package seeder

import "github.com/harisfi/alterra-agmc/day10/database"

func Seed() {
	conn := database.GetConnection()

	userSeeder(conn)
	bookSeeder(conn)
}
