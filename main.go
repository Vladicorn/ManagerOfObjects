package main

import (
	"TaskManager/database"
	"TaskManager/handlers"
	"taskmanager/telegram"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	key := "1915374423:AAF0QOTMcHUjTksdwkLLb3dYOiqDaxlLbCU"
	database.Connect()
	//	fmt.Println(database.Select())
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*",
		AllowHeaders:     "",
	}))
	handlers.Registor(app)
	go telegram.ConnectTG(key)
	app.Listen(":8000")

}
