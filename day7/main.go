package main

import (
	"os"

	"github.com/harisfi/alterra-agmc/day7/database"
	// "github.com/harisfi/alterra-agmc/day7/database/migration"
	// "github.com/harisfi/alterra-agmc/day7/database/seeder"
	"github.com/harisfi/alterra-agmc/day7/internal/factory"
	"github.com/harisfi/alterra-agmc/day7/internal/http"
	"github.com/harisfi/alterra-agmc/day7/internal/middleware"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	if e := godotenv.Load(); e != nil {
		panic(e)
	}
}

func main() {
	database.CreateConnection()
	// migration.Migrate()
	// seeder.Seed()

	f := factory.NewFactory()
	e := echo.New()
	middleware.Init(e)
	http.NewHttp(e, f)

	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}
