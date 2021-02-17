package apps

import (
	"net/http"

	db "coffee-mate/src/database"
	"coffee-mate/src/database/migration"

	"coffee-mate/src/middleware/exception"
	"coffee-mate/src/routes"

	"github.com/gin-gonic/gin"
)

// Application -> application instance
type Application struct {
}

// CreateApp -> method to create gin application
func (a Application) CreateApp(r *gin.Engine) {
	// r.Use(exception.Recovery(exception.ErrorHandler))
	configureAppDB()
	configureAPIEndpoint(r)
}

// CreateTest -> method to create gin application with environment test
func (a Application) CreateTest(r *gin.Engine) {
	r.Use(exception.Recovery(exception.ErrorHandler))
	configureTestDB()
	configureAPIEndpoint(r)
}

/**
 * configuration all endpoint
 */
func configureAPIEndpoint(r *gin.Engine) {
	g := r.Group("/api/v1")
	routes.Router(g)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"flag":   "NOT_FOUND",
			"errors": gin.H{
				"message": "The route you are looking for is not found",
				"flag":    "ROUTE_NOT_FOUND",
			},
		})
	})
}

/**
 * configuration database application
 */
func configureAppDB() {
	db.AppConnection()
}

/**
 * configuration database application for testing
 */
func configureTestDB() {
	db.TestConnection()
	conn := db.GetDB()
	migration.AutoMigration(conn)
}
