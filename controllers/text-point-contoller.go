package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"schedule/models"
)

func UpdateTextPoint(c *gin.Context) {

	var input = models.TextPoint{}

	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println("1234")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.UpdateTextPoint()

	c.JSON(http.StatusOK, gin.H{"status": "success"})

}
