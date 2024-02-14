package web

type UserCreateResponse struct {
	Username string `json:"username,omitempty"`
	Message  string `json:"message,omitempty"`
}
