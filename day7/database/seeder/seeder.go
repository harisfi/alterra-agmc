package seeder

import "github.com/harisfi/alterra-agmc/day7/database"

func Seed() {
	conn := database.GetConnection()

	userSeeder(conn)
	bookSeeder(conn)
}
