package routes

import (
	"net/http"

	"coffee-mate/src/controllers"
	"coffee-mate/src/middleware/auth"
	"coffee-mate/src/validations"

	"github.com/gin-gonic/gin"
)

// Router middleware to handler routes
func Router(g *gin.RouterGroup) {
	usercontroller := controllers.UController()
	// homeController := controllers.HController()
	// {
	// 	g.GET("/users", controller.GetUsers)
	// 	g.GET("/user/:id", validations.GetUser, controller.GetUser)
	g.POST("/signup", validations.CreateUser, usercontroller.CreateUser)
	// 	g.PATCH("/user/:id", validations.UpdateUser, controller.UpdateUser)
	// 	g.DELETE("/user/:id", validations.DeleteUser, controller.DeleteUser)

	g.POST("/login", validations.LoginUser, usercontroller.LoginUser)

	// home
	g.GET("/", auth.TokenAuthenticationMiddleware, controllers.HomeMessage)
	g.GET("/home", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Welcome to Coffee-Mate Application, The home of all coffee",
		})
	})
}
