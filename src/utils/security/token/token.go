package token

import (
	"coffee-mate/src/middleware/exception"
	"coffee-mate/src/services/redis"
	"fmt"

	"coffee-mate/src/validations/schemas"

	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

var red = redis.RService()

// CreateToken -> Create Token Method
func CreateToken(id uuid.UUID, email, firstName string) (string, error) {
	config := red.Load()
	var tokenString string
	var err error
	expiryTime := config["token-expiration"]
	if expiryTime == "" {
		expiryTime = "1800"
	}
	// convert expiryTime to integer
	i, _ := strconv.ParseInt(expiryTime, 10, 32)
	// golang time is not integer, so convert int to time
	// yourTime := rand.Int63n(i)

	var n time.Duration = time.Duration(i) * time.Second

	expiresAt := time.Now().Add(n).Unix()

	// check if the user has a token in redis
	authToken, err := red.RedisHGet(red.LoginTokenStore, id.String())
	// expiresAt := time.Now().Add(time.Minute * n).Unix()
	atClaims := jwt.MapClaims{}
	//Creating Access Claims
	atClaims["authorized"] = true
	atClaims["UserID"] = id
	atClaims["FirstName"] = firstName
	atClaims["Email"] = email
	atClaims["exp"] = expiresAt

	if authToken != "" || err == nil {
		// check if the token has not expired
		expired := CheckTokenExpiryAndValidity(authToken)
		if expired {
			//Creating Access Token
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
			tokenString, err = token.SignedString([]byte(os.Getenv("API_SECRET")))
			// Save to redis
			rData := schemas.RedisData{
				ID:    red.LoginTokenStore,
				Key:   id.String(),
				Value: tokenString,
			}
			red.RedisHSet(rData)
			return tokenString, err
		}
		return authToken, nil
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	tokenString, err = token.SignedString([]byte(os.Getenv("API_SECRET")))
	// Save to redis
	rData := schemas.RedisData{
		ID:    red.LoginTokenStore,
		Key:   id.String(),
		Value: tokenString,
	}
	red.RedisHSet(rData)
	return tokenString, err
}

// VerifyToken -> validation for login user
func VerifyToken(c *gin.Context) (*jwt.Token, error) {
	var errors []map[string]interface{}
	tokenString := ExtractToken(c)

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

// CheckTokenExpiryAndValidity ->
func CheckTokenExpiryAndValidity(token string) bool {
	authToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("API_SECRET")), nil
	})

	v, _ := err.(*jwt.ValidationError)
	if !authToken.Valid {
		return false
	}

	if v != nil {
		if v.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			return false
		}
		return false
	}

	return true
}
