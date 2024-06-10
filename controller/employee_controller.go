package controllers

import (
	"hotel-management/models"
	"hotel-management/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	EmployeeService services.EmployeeService
}

func NewEmployeeController(employeeService services.EmployeeService) *EmployeeController {
	return &EmployeeController{
		EmployeeService: employeeService,
	}
}

func (ec *EmployeeController) CreateEmployee(c *gin.Context) {
	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ec.EmployeeService.CreateEmployee(&employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, employee)
}

func (ec *EmployeeController) GetAllEmployees(c *gin.Context) {
	employees, err := ec.EmployeeService.GetAllEmployees()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, employees)
}
