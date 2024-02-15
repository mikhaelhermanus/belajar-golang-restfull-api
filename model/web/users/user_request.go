package web

type User struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
}
