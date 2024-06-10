package repositories

import (
	"hotel-management/models"
	"hotel-management/utils"
)

type EmployeeRepository interface {
	CreateEmployee(employee *models.Employee) error
	GetAllEmployees() ([]models.Employee, error)
}

type employeeRepository struct{}

func NewEmployeeRepository() EmployeeRepository {
	return &employeeRepository{}
}

func (er *employeeRepository) CreateEmployee(employee *models.Employee) error {
	return utils.DB.Create(employee).Error
}

func (er *employeeRepository) GetAllEmployees() ([]models.Employee, error) {
	var employees []models.Employee
	err := utils.DB.Find(&employees).Error
	return employees, err
}
