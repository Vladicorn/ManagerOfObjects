package handlers

import (
	"taskmanager/controllers"
	"taskmanager/middleware"

	"github.com/gofiber/fiber/v2"
)

func Registor(app *fiber.App) {

	//	app.Get("/api/users", controllers.AllTask)

	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Use(middleware.AuthMiddleware)

	app.Get("/api/logout", controllers.Logout)
	app.Get("/api/users", controllers.AllUsers)
	app.Post("/api/users", controllers.CreaterUser)

	app.Get("/api/objects", controllers.AllObject)
	app.Get("/api/object/:id", controllers.GetObj)
	app.Post("/api/object", controllers.CreateObj)
	app.Put("/api/object/:id", controllers.UpdateObj)
	app.Delete("/api/object/:id", controllers.DeleteObj)

	//app.Get("/api/object/:id/task/:num", controllers.GetObjWithCom)

}
