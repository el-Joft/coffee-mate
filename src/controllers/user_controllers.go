package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
)

// UserController -> the propose of user controller
// is handling request from client and
// forward it to specific service
type UserController struct {}

// CreateUser -> create user routes
// POST /user
func (u *UserController) CreateUser(c *gin.Context)  {
	log.Println("===================")
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": flag.GetUsersSuccess.Message,
		"data":    users,
		"error":   nil,
	})
}