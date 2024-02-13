package web

type UserCreateResponse struct {
	Username string `json:"username"`
	Message  string `json:"message,omitempty"`
}
