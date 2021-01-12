package tests

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"coffee-mate/src/apps"
	db "coffee-mate/src/database"
	"coffee-mate/src/utils/response"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/stretchr/testify/assert"
)

func initTestCreateUser(body map[string]interface{}) (*httptest.ResponseRecorder, *gin.Engine) {
	r := gin.Default()
	app := new(apps.Application)
	app.CreateTest(r)

	w := httptest.NewRecorder()
	b, _ := json.Marshal(body)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/signup", strings.NewReader(string(b)))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)
	return w, r
}

// TestCreateUser -> Testing User creatio
func TestCreateUser(t *testing.T) {
	t.Run("it should return success", func(t *testing.T) {
		defer db.DropAllTable()
		body := map[string]interface{}{
			"email":      "test@gmail.com",
			"first_name": "Test",
			"last_name":  "TestLastName",
			"username":   "username",
			"age":        20,
			"password":   "1231212",
		}
		w, _ := initTestCreateUser(body)

		actual := response.Response{}
		if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
			panic(err)
		}
		log.Printf("Data %s\n", w.Body.Bytes())
		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Equal(t, http.StatusCreated, actual.Status)
		assert.NotEmpty(t, actual.Data)
		assert.Empty(t, actual.Errors)
	})

	t.Run("it should return user already exist", func(t *testing.T) {
		defer db.DropAllTable()
		body := map[string]interface{}{
			"email":      "test@gmail.com",
			"first_name": "Test",
			"last_name":  "TestLastName",
			"username":   "usernamer",
			"age":        20,
			"password":   "1231212",
		}
		_, r := initTestCreateUser(body)

		w := httptest.NewRecorder()
		b, _ := json.Marshal(body)

		req, _ := http.NewRequest(http.MethodPost, "/api/v1/signup", strings.NewReader(string(b)))
		req.Header.Set("Content-Type", "application/json")

		r.ServeHTTP(w, req)

		actual := response.Response{}
		if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
			panic(err)
		}

		assert.Equal(t, http.StatusConflict, w.Code)
		assert.Equal(t, http.StatusConflict, actual.Status)
		assert.Equal(t, "User Conflict", actual.Message)
		assert.Equal(t, "User with this email already exist", actual.Errors[0].Message)
		assert.Equal(t, "USER_ALREADY_EXIST", actual.Errors[0].Flag)
	})

	t.Run("it should return invalid body with invalid password format", func(t *testing.T) {
		defer db.DropAllTable()
		body := map[string]interface{}{
			"email":      "testy@gmail.com",
			"first_name": "Test",
			"last_name":  "TestLastName",
			"username":   "username",
			"age":        20,
			"password":   "123",
		}
		w, _ := initTestCreateUser(body)

		// log.Printf("Data %s\n", w.Body.Bytes())

		actual := response.Response{}
		if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
			panic(err)
		}
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, http.StatusBadRequest, actual.Status)
		assert.Equal(t, "Validation error", actual.Message)
		assert.Equal(t, "INVALID_BODY", actual.Errors[0].Flag)
		assert.NotEmpty(t, actual.Errors[0].Message)
	})

	t.Run("it should return invalid body with invalid email format", func(t *testing.T) {
		defer db.DropAllTable()
		body := map[string]interface{}{
			"email":      "test.com",
			"first_name": "FristName",
			"last_name":  "LastName",
			"username":   "testusername",
			"age":        20,
			"password":   "123121ABC",
		}
		w, _ := initTestCreateUser(body)

		actual := response.Response{}
		if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
			panic(err)
		}

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, http.StatusBadRequest, actual.Status)
		assert.Equal(t, "Validation error", actual.Message)
		assert.Equal(t, "INVALID_BODY", actual.Errors[0].Flag)
		assert.NotEmpty(t, actual.Errors[0].Message)
	})

}
