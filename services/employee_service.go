package services

import (
	"hotel-management/models"
	"hotel-management/repositories"
)

type EmployeeService interface {
	CreateEmployee(employee *models.Employee) error
	GetAllEmployees() ([]models.Employee, error)
}

type employeeService struct {
	EmployeeRepository repositories.EmployeeRepository
}

func NewEmployeeService(employeeRepo repositories.EmployeeRepository) EmployeeService {
	return &employeeService{
		EmployeeRepository: employeeRepo,
	}
}

func (es *employeeService) CreateEmployee(employee *models.Employee) error {
	return es.EmployeeRepository.CreateEmployee(employee)
}

func (es *employeeService) GetAllEmployees() ([]models.Employee, error) {
	return es.EmployeeRepository.GetAllEmployees()
}
