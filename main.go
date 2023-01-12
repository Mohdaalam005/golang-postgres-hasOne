package main

import (
	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
	"github.com/mohdaalam005/one-to-one/database"
	"github.com/mohdaalam005/one-to-one/route"
	"log"
)

func main() {
	app := fiber.New()
	log.Println("Your application has been started !")
	godotenv.Load("config.env")
	database.Connect()
	route.Route(app)

}
