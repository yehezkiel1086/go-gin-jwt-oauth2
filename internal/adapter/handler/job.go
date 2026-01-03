package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-gin-nextjs-auth/internal/core/domain"
	"github.com/yehezkiel1086/go-gin-nextjs-auth/internal/core/port"
)

type JobHandler struct {
	svc port.JobService
}

func NewJobHandler(svc port.JobService) *JobHandler {
	return &JobHandler{
		svc,
	}
}

type CreateJobReq struct {
	Title string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Company string `json:"company" binding:"required"`
	Location string `json:"location" binding:"required"`
	Salary float64 `json:"salary" binding:"required"`
}

func (jh *JobHandler) CreateJob(c *gin.Context) {
	var req CreateJobReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	job, err := jh.svc.CreateJob(c, &domain.Job{
		Title: req.Title,
		Description: req.Description,
		Company: req.Company,
		Location: req.Location,
		Salary: req.Salary,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, job)
}

func (jh *JobHandler) GetJobs(c *gin.Context) {
	jobs, err := jh.svc.GetJobs(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": domain.ErrNotFound,
		})
		return
	}

	c.JSON(http.StatusOK, jobs)
}

func (jh *JobHandler) GetJobById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		errMsg := fmt.Errorf("invalid id parameter: %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errMsg,
		})
		return
	}

	job, err := jh.svc.GetJobById(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, job)
}

func (jh *JobHandler) DeleteJob(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		errMsg := fmt.Errorf("invalid id parameter: %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errMsg,
		})
		return		
	}

	if err := jh.svc.DeleteJob(c.Request.Context(), uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": domain.ErrInternal,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "job is deleted successfully",
	})
}
