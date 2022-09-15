package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/harisfi/alterra-agmc/day3/configs"
	"github.com/harisfi/alterra-agmc/day3/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	configs.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8080"))
}
