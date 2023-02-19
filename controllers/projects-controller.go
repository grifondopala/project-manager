package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"schedule/models"
	"strconv"
)

type CreateProjectInput struct {
	UserID uint `json:"user_id" binding:"required"`
}

func CreateProject(c *gin.Context) {

	var input = CreateProjectInput{}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project, err := models.CreateProject(input.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	defaultProjectTemplate(project.ID)

	c.JSON(http.StatusOK, gin.H{"message": "created project success", "data": project})
}

func defaultProjectTemplate(pId uint) {
	NameCol, _ := (&models.Column{
		Name:        "Task name",
		OrderNumber: 1,
		ProjectID:   pId,
		Type:        "text-point",
	}).Create()
	DescriptionCol, _ := (&models.Column{
		Name:        "Task description",
		OrderNumber: 2,
		ProjectID:   pId,
		Type:        "text-point",
	}).Create()
	TaskFirst, _ := (&models.Task{
		OrderNumber: 1,
		ProjectID:   pId,
		Section:     "First section",
	}).Create()
	_, err := (&models.TextPoint{
		Text:     "Make your list",
		TaskID:   TaskFirst.ID,
		ColumnID: NameCol.ID,
	}).Create()
	if err != nil {
		log.Fatal(err)
	}
	_, err = (&models.TextPoint{
		Text:     "This is interesting task!",
		TaskID:   TaskFirst.ID,
		ColumnID: DescriptionCol.ID,
	}).Create()
	if err != nil {
		log.Fatal(err)
	}
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

	Columns, Tasks := GetProjectTasks(project.ID)

	c.JSON(http.StatusOK, gin.H{"projectInfo": project, "tasks": Tasks, "columns": Columns})
}

type TaskT struct {
	Task   models.Task   `json:"task"`
	Points []interface{} `json:"points"`
}

func GetProjectTasks(pId uint) ([]*models.Column, []TaskT) {
	Columns, errC := (&models.Column{}).GetProjectColumns(pId)
	if errC != nil {

	}
	Tasks, errT := (&models.Task{}).GetProjectTasks(pId)
	if errT != nil {

	}
	var arrTasks []TaskT
	for _, task := range Tasks {
		var newTask TaskT
		newTask.Task = *task
		for _, column := range Columns {
			switch column.Type {
			case "text-point":
				{
					point, err := (&models.TextPoint{}).GetTextPoint(task.ID, column.ID)
					if err != nil {

					}
					newTask.Points = append(newTask.Points, point)
				}
			default:
				fmt.Println("Not fount")
			}
		}
		arrTasks = append(arrTasks, newTask)
	}
	return Columns, arrTasks
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
