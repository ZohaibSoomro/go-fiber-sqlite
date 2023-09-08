package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Name       string `json:"name"`
	Department string `json:"department"`
}

func NewEmployee(name string, department string) *Employee {
	return &Employee{Name: name, Department: department}
}

// const (
// 	OPERATIONS     department = "Operations"
// 	DEVELOPMENT    department = "Development"
// 	FINANCE        department = "Finance"
// 	ADMINISTRATION department = "Administration"
// )
