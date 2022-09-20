package main

import "github.com/joho/godotenv"

func init() {
	if e := godotenv.Load(); e != nil {
		panic(e)
	}
}

func main() {
	//
}
