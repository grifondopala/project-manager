package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"schedule/models"
)

func CreateColumn(c *gin.Context) {

	var column = models.Column{}
	if err := c.ShouldBindJSON(&column); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	col, err := column.Create()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Tasks, _ := (&models.Task{}).GetProjectTasks(column.ProjectID)

	var Points []interface{}

	for _, task := range Tasks {
		switch column.Type {
		case "text-point":
			{
				point, _ := (&models.TextPoint{Text: "", ColumnID: col.ID, TaskID: (*task).ID}).Create()
				Points = append(Points, point)
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"column": col, "points": Points})

}

func UpdateColumn(c *gin.Context) {

	var input = models.UpdateColumInput{}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := models.UpdateColumn(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})

}
