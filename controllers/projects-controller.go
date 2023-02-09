package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"schedule/models"
	"strconv"
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

func GetUserProjects(c *gin.Context) {

	userId, err := strconv.Atoi(c.Param("user_id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	projects, err := models.GetUserProjects(uint(userId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": projects})

}

func GetProjectById(c *gin.Context) {

	projectId, err := strconv.Atoi(c.Param("project_id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project, err := models.GetProjectById(uint(projectId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "projectInfo": project})

}

func UpdateInformation(c *gin.Context) {

	var input = models.UpdateProjectInput{}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := models.UpdateInformation(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})

}
