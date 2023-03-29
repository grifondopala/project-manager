package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"schedule/models"
)

func UpdateTextPoint(c *gin.Context) {

	var input = models.TextPoint{}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.UpdateTextPoint()

	c.JSON(http.StatusOK, gin.H{"status": "success"})

}
