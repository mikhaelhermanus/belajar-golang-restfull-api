package web

type CategoryResponse struct {
	// if isempty message , will return undefined
	Id      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Message string `json:"message,omitempty"`
}
