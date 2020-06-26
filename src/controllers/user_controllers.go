package controllers

import (
	"coffee-mate/src/database/entity"
	"coffee-mate/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// UserController -> the propose of user controller
// is handling request from client and
// forward it to specific service
type UserController struct {
	Service services.UserService
}

// UController -> user controller instance
func UController() UserController {
	return UserController{
		Service: services.UService(),
	}
}

// CreateUser -> create user routes
// POST /user
func (u *UserController) CreateUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindBodyWith(&user, binding.JSON); err != nil {
		c.AbortWithStatus(400)
		return
	}
	// log.Printf("Data %s\n", user)
	data := u.Service.CreateUser(user)
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Success create user",
		"data":    data,
		"error":   nil,
	})
}
