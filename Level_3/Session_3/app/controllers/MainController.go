package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":       "Gin Web Apps V1.0",
		"description": "Sample Web Apps with GIN in GO",
	})
}

// Get Parameter from Body
func PostResult(c *gin.Context) {
	input1 := c.PostForm("input1A")
	input2 := c.PostForm("input2A")

	c.HTML(http.StatusOK, "result.html", gin.H{
		"title":  "Gin Web Apps V1.0 (Form Result)",
		"input1": input1,
		"input2": input2,
	})
}

// Get Parameter from URL Query
func GetResultQuery(c *gin.Context) {
	input1 := c.Query("input1B")
	input2 := c.Query("input2B")

	c.HTML(http.StatusOK, "result.html", gin.H{
		"title":  "Gin Web Apps V1.0 (Form Result)",
		"input1": input1,
		"input2": input2,
	})
}

// Get Parameter from URL Without Query
func GetResult(c *gin.Context) {
	input1 := c.Param("input1")
	input2 := c.Param("input2")

	c.HTML(http.StatusOK, "result.html", gin.H{
		"title":  "Gin Web Apps V1.0 (Form Result)",
		"input1": input1,
		"input2": input2,
	})
}
