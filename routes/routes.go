package routes

import (
	"github.com/gofiber/fiber"
	"github.com/zohaibsoomro/go-fiber-sqlite/controllers"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/employees", controllers.GetEmployees)
	app.Get("/employees/:id", controllers.GetEmployeeWithId)
	app.Post("/employee/create", controllers.CreateEmployee)
	app.Put("/employees/:id", controllers.UpdateEmployee)
	app.Delete("/employees/:id", controllers.DeleteEmployee)
}
