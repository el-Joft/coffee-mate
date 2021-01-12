package security

import (
	"coffee-mate/src/validations/schemas"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

// CreateToken -> Create Token Method
func CreateToken(id uuid.UUID, email, firstName string) (string, error) {

	// Todo: the expiry time should be from redis
	expiresAt := time.Now().Add(time.Minute * 100000).Unix()

	tk := &schemas.Token{
		UserID: id,
		Name:   firstName,
		Email:  email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tk)
	tokenString, error := token.SignedString([]byte(os.Getenv("API_SECRET")))

	return tokenString, error
}
