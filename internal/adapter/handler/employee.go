package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/core/domain"
	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/core/port"
)

type EmployeeHandler struct {
	svc port.EmployeeService
}

func InitEmployeeHandler(svc port.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{
		svc: svc,
	}
}

type CreateEmployeeReq struct {
	Name string `json:"name" binding:"required"`
	Position string `json:"position" binding:"required"`
	Description string `json:"description"`
}

func (eh *EmployeeHandler) CreateEmployee(c *gin.Context) {
	// bind input
	var input *CreateEmployeeReq
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "name and position are required",
		})
		return
	}

	// create employee
	res, err := eh.svc.CreateEmployee(c, &domain.Employee{
		Name: input.Name,		
		Position: input.Position,
		Description: input.Description,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// return response
	c.JSON(http.StatusCreated, res)
}

func (eh *EmployeeHandler) GetEmployees(c *gin.Context) {
	emps, err := eh.svc.GetEmployees(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to retrieve employees",
		})
		return
	}

	c.JSON(http.StatusOK, emps)
}
