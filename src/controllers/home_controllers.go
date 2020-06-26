package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HomeController -> the propose of Home controller
// is handling request from client and
// forward it to specific service

// HomeMessage -> get home page routes
// GET /
func HomeMessage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Welcome to Coffee-Mate Application, The home of all coffee",
		"data":    nil,
		"error":   nil,
	})
}
