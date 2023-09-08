package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber"
	"github.com/zohaibsoomro/go-fiber-sqlite/db"
	"github.com/zohaibsoomro/go-fiber-sqlite/models"
)

var CreateEmployee = func(ctx *fiber.Ctx) {
	employee := new(models.Employee)
	ctx.BodyParser(employee)
	d := db.GetDb().Create(employee)
	if d.Error != nil {
		ctx.Status(http.StatusExpectationFailed).SendString(fmt.Sprintf("creation failed: %s", d.Error))
		return
	}
	ctx.SendString("Employee created.")
	ctx.Status(http.StatusOK).JSON(employee)
}

var GetEmployeeWithId = func(ctx *fiber.Ctx) {
	i := ctx.Params("id")
	id, err := strconv.Atoi(i)
	if err != nil {
		ctx.Status(http.StatusBadRequest).SendString("Invalid id")
		return
	}
	employee := new(models.Employee)
	d := db.GetDb().Where("Id = ?", id).First(employee)
	if d.Error != nil {
		ctx.Status(http.StatusExpectationFailed).SendString(fmt.Sprintf("error occurred: %s", d.Error))
		return
	}
	ctx.Status(http.StatusOK).JSON(employee)
}
var GetEmployees = func(ctx *fiber.Ctx) {
	employees := make([]models.Employee, 0)
	d := db.GetDb().Find(&employees)
	if d.Error != nil {
		ctx.Status(http.StatusExpectationFailed).SendString(fmt.Sprintf("error occurred: %s", d.Error))
		return
	}
	ctx.Status(http.StatusOK).JSON(employees)
}
var UpdateEmployee = func(ctx *fiber.Ctx) {
	i := ctx.Params("id")
	id, err := strconv.Atoi(i)
	if err != nil {
		ctx.Status(http.StatusBadRequest).SendString("Invalid id.")
		return
	}
	employeeRecord := new(models.Employee)
	db.GetDb().Where("Id = ?", id).First(employeeRecord)
	//if record not found
	if employeeRecord.Name == "" {
		ctx.Status(http.StatusNotFound).SendString("employee not found.")
		return
	}
	employee := new(models.Employee)
	err = ctx.BodyParser(employee)
	if err != nil {
		ctx.Status(http.StatusNotAcceptable).SendString(fmt.Sprintf("error while parsing body: %s", err))
		return
	}
	if employee.Name != "" {
		employeeRecord.Name = employee.Name
	}
	if employee.Department != "" {
		employeeRecord.Department = employee.Department
	}
	d := db.GetDb().Save(employeeRecord)

	if d.Error != nil {
		ctx.Status(http.StatusExpectationFailed).SendString(fmt.Sprintf("error occurred %s", d.Error))
		return
	}
	ctx.Status(http.StatusOK).JSON(employeeRecord)
}
var DeleteEmployee = func(ctx *fiber.Ctx) {
	i := ctx.Params("id")
	id, err := strconv.Atoi(i)
	if err != nil {
		ctx.Status(http.StatusBadRequest).SendString("Invalid id.")
		return
	}
	employeeRecord := new(models.Employee)
	db.GetDb().Where("Id = ?", id).First(employeeRecord)
	//if record not found
	if employeeRecord.Name == "" {
		ctx.Status(http.StatusNotFound).SendString("employee not found.")
		return
	}
	d := db.GetDb().Where("id=?", id).Delete(employeeRecord)
	if d.Error != nil {
		ctx.Status(http.StatusExpectationFailed).SendString(fmt.Sprintf("error occurred %s", d.Error))
		return
	}
	ctx.Status(http.StatusOK).SendString("employee deleted.")
	ctx.JSON(employeeRecord)
}
