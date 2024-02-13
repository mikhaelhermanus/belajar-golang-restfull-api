package web

type UserCreateRequest struct {
	Email    string `validate:"required,min=1,max=100" json:"email"`
	Phone    string `validate:"required,min=1,max=100" json:"phone"`
	Username string `validate:"required,min=1,max=100" json:"username"`
	Password string `validate:"required,min=1,max=100" json:"password"`
	Name     string `validate:"required,min=1,max=100" json:"name"`
}
