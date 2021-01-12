package schemas

// CreateUser is create user schema validation
type CreateUser struct {
	FirstName string `validate:"required" json:"first_name"`
	LastName  string `validate:"required" json:"last_name"`
	Username  string `validate:"required"`
	Email     string `validate:"required,email"`
	Password  string `json:"password" validate:"required,passwd"`
	Age       int64  `validate:"omitempty,numeric,gt=0"`
}

// UserID is param uri validation for id
type UserID struct {
	ID uint `uri:"id" binding:"required"`
}

// UpdateUser is update user schema validation
type UpdateUser struct {
	Name    string
	Email   string `validate:"omitempty,email"`
	Age     int64  `validate:"omitempty,numeric,gt=0"`
	Address string
}

// LoginUser is login user schema validation
type LoginUser struct {
	Email    string `validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
