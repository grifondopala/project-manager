package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"schedule/models"
)

func CreateEmptyTask(c *gin.Context) {

	var input = models.Task{}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Task, _ := (&input).Create()

	Columns, _ := (&models.Column{}).GetProjectColumns(input.ProjectID)

	var Points []interface{}

	for _, column := range Columns {
		switch column.Type {
		case "text-point":
			{
				point, _ := (&models.TextPoint{Text: "", ColumnID: column.ID, TaskID: (*Task).ID}).Create()
				Points = append(Points, point)
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"task": Task, "points": Points})

}

func UpdateTask(c *gin.Context) {

	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := task.Update()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})

}
