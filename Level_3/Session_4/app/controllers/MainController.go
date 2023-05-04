package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Show multiple template (index)
func ShowIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index", gin.H{
		"title": "Index",
		"year":  time.Now().Year(),
	})
}

// Show multiple template (about)
func ShowAbout(c *gin.Context) {
	c.HTML(http.StatusOK, "about", gin.H{
		"title": "About",
		"year":  time.Now().Year(),
	})
}

// Show multiple template (contact)
func ShowContact(c *gin.Context) {
	c.HTML(http.StatusOK, "contact", gin.H{
		"title": "Contact",
		"year":  time.Now().Year(),
	})
}
