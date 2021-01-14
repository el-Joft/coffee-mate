package controllers

import (
	// "coffee-mate/src/validations/schemas"
	"coffee-mate/src/database/entity"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HomeController -> the propose of Home controller
// is handling request from client and
// forward it to specific service

// HomeMessage -> get home page routes
// GET /
func HomeMessage(c *gin.Context) {
	user := c.MustGet("user").(entity.User)
	// log.Println(ok)
	log.Println(user.Email)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Welcome to Coffee-Mate Application, The home of all coffee",
		"data":    nil,
		"error":   nil,
	})
}
