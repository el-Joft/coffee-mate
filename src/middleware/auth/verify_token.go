package auth

import (
	"coffee-mate/src/database/entity"
	"coffee-mate/src/middleware/exception"
	"coffee-mate/src/repositories"
	"coffee-mate/src/utils/security"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// TokenAuthenticationMiddleware ->
func TokenAuthenticationMiddleware(c *gin.Context) {
	token := TokenValid(c)
	loggedInUser := FetchAuth(token)
	c.Set("user", loggedInUser)
	c.Next()

}

// TokenValid ->
func TokenValid(c *gin.Context) *jwt.Token {
	var errors []map[string]interface{}
	token, err := security.VerifyToken(c)

	if err != nil {
		errors = append(errors, map[string]interface{}{
			"message": err.Error()},
		)
		exception.StatusUnauthorized("error", errors)
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		errors = append(errors, map[string]interface{}{
			"message": err.Error()},
		)
		exception.StatusUnauthorized("error", errors)
	}
	return token
}

// FetchAuth ->
func FetchAuth(token *jwt.Token) entity.User {
	var errors []map[string]interface{}
	// get claim from token
	claims, _ := token.Claims.(jwt.MapClaims)
	UserID, _ := claims["UserID"].(string)
	r := repositories.URepository()
	user, err := r.GetByParam(UserID)
	if err != nil {
		errors = append(errors, map[string]interface{}{
			"message": "Invalid user"},
		)
		exception.StatusUnauthorized("error", errors)
	}
	return user

}
