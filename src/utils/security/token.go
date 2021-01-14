package security

import (
	"coffee-mate/src/middleware/exception"
	"fmt"

	// "coffee-mate/src/validations/schemas"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// CreateToken -> Create Token Method
func CreateToken(id uuid.UUID, email, firstName string) (string, error) {

	// Todo: the expiry time should be from redis
	expiresAt := time.Now().Add(time.Minute * 100000).Unix()

	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["UserID"] = id
	atClaims["FirstName"] = firstName
	atClaims["Email"] = email
	atClaims["exp"] = expiresAt

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	tokenString, error := token.SignedString([]byte(os.Getenv("API_SECRET")))

	return tokenString, error
}

// VerifyToken -> validation for login user
func VerifyToken(c *gin.Context) (*jwt.Token, error) {
	var errors []map[string]interface{}
	tokenString := ExtractToken(c)

	// tk := &schemas.Token{}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		errors = append(errors, map[string]interface{}{
			"message": err.Error()},
		)
		exception.StatusUnauthorized("error", errors)
	}

	return token, nil

}

// ExtractToken ->
func ExtractToken(c *gin.Context) string {
	var errors []map[string]interface{}
	bearToken := c.Request.Header.Get("Authorization")
	bearToken = strings.TrimSpace(bearToken)

	if bearToken == "" {
		//Token is missing, returns with error code 403 Unauthorized
		errors = append(errors, map[string]interface{}{
			"message": "Authentication Token is missing"},
		)
		exception.StatusForbidden("error", errors)
	}

	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")

	if len(strArr) == 2 && strArr[0] == "Bearer" {
		return strArr[1]
	}
	return ""
}
