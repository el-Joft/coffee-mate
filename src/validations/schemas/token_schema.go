package schemas

import (
	jwt "github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

//Token struct declaration
type Token struct {
	UserID uuid.UUID
	Name   string
	Email  string
	*jwt.StandardClaims
}
