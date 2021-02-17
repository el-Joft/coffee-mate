package controllers

import (
	"coffee-mate/src/database/entity"
	// "coffee-mate/src/services"

	// "coffee-mate/src/validations/schemas"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HomeController -> the propose of Home controller
// is handling request from client and
// forward it to specific service
type HomeController struct {
	// RedisService services.RedisService
}

// HController -> user controller instance
func HController() HomeController {
	return HomeController{
		// RedisService: services.RService(),
	}
}

// HomeMessage -> get home page routes
// GET /
func (h *HomeController) HomeMessage(c *gin.Context) {
	user := c.MustGet("user").(entity.User)

	// sare set redis
	// rData := schemas.RedisData{
	// 	ID:    "REDIS_APP_CONFIG_KEY",
	// 	Key:   "token-expiration",
	// 	Value: "10",
	// }
	log.Println(user)

	// r2Data := schemas.RedisData{
	// 	ID:    "REDIS_APP_CONFIG_KEY",
	// 	Key:   user.ID.String(),
	// 	Value: "Nsdsdsds",
	// }
	// h.RedisService.RedisHSet(rData)
	// h.RedisService.RedisHSet(r2Data)

	log.Println("=====================")

	// result := h.RedisService.Load()
	// log.Println(result)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Welcome to Coffee-Mate Application, The home of all coffee",
		"data":    nil,
		"error":   nil,
	})
}
