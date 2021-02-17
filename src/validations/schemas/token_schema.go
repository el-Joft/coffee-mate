package schemas

import (
	"encoding/json"

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

// RedisData ->
type RedisData struct {
	ID    string
	Key   string
	Value string
}

// MarshalBinary ->
func (t *RedisData) MarshalBinary() ([]byte, error) {
	return json.Marshal(t)
}

//UnmarshalBinary ->
func (t *RedisData) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}

	return nil
}
