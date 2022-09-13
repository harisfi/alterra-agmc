package main

import (
	"github.com/harisfi/alterra-agmc/day2/learn-mvc/configs"
	"github.com/harisfi/alterra-agmc/day2/learn-mvc/routes"
)

func main() {
	configs.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
