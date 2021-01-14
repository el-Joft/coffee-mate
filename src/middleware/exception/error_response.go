package exception

import "net/http"

// NotFound -> response empty data
func NotFound(message string, errors []map[string]interface{}) {
	response := map[string]interface{}{
		"status":  http.StatusNotFound,
		"message": message,
		"data":    nil,
		"errors":  errors,
	}
	panic(response)
}

// BadRequest -> response for bad request
func BadRequest(message string, errors []map[string]interface{}) {
	response := map[string]interface{}{
		"status":  http.StatusBadRequest,
		"message": message,
		"data":    nil,
		"errors":  errors,
	}
	panic(response)
}

// Conflict -> response for conflict
func Conflict(message string, errors []map[string]interface{}) {
	response := map[string]interface{}{
		"status":  http.StatusConflict,
		"message": message,
		"data":    nil,
		"errors":  errors,
	}
	panic(response)
}

// InternalServerError -> response for internal server error
func InternalServerError(message string, errors []map[string]interface{}) {
	response := map[string]interface{}{
		"message": message,
		"data":    nil,
		"status":  http.StatusInternalServerError,
		"errors":  errors,
	}
	panic(response)
}

// StatusUnauthorized -> response for internal server error
func StatusUnauthorized(message string, errors []map[string]interface{}) {
	response := map[string]interface{}{
		"message": message,
		"data":    nil,
		"status":  http.StatusUnauthorized,
		"errors":  errors,
	}
	panic(response)
}

// StatusUnprocessableEntity -> response for unprocessable entity
func StatusUnprocessableEntity(message string, errors []map[string]interface{}) {
	response := map[string]interface{}{
		"message": message,
		"data":    nil,
		"status":  http.StatusUnprocessableEntity,
		"errors":  errors,
	}
	panic(response)
}

// StatusForbidden -> response for unprocessable entity
func StatusForbidden(message string, errors []map[string]interface{}) {
	response := map[string]interface{}{
		"message": message,
		"data":    nil,
		"status":  http.StatusForbidden,
		"errors":  errors,
	}
	panic(response)
}
