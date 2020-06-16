package schemas

// CreateUser is create user schema validation
type CreateUser struct {
	Fullname string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `json:"password" validate:"required,min=4,passwd"`
	Age      int64  `validate:"omitempty,numeric,gt=0"`
	Address  string
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
