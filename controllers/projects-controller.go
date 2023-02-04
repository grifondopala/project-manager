package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"schedule/models"
)

type CreateProjectInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	IconSrc     string `json:"icon_src" binding:"required"`
	Color       string `json:"color" binding:"required"`
	UserID      uint   `json:"user_id" binding:"required"`
}

func CreateProject(c *gin.Context) {

	var input = CreateProjectInput{}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project := models.Project{}

	project.Name = input.Name
	project.Description = input.Description
	project.IconSrc = input.IconSrc
	project.UserID = input.UserID
	project.Color = input.Color

	project, err := project.SaveProject()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "created project success", "data": project})

}

type GetUserProjectsInput struct {
	UserID uint `json:"user_id" binding:"required"`
}

func GetUserProjects(c *gin.Context) {

	var input = GetUserProjectsInput{}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	projects, err := models.GetUserProjects(input.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": projects})

}

type GetProjectByIdInput struct {
	ProjectID uint `json:"project_id" binding:"required"`
}

func GetProjectById(c *gin.Context) {

	var input = GetProjectByIdInput{}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(input.ProjectID)

	project, err := models.GetProjectById(input.ProjectID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "projectInfo": project})

}
