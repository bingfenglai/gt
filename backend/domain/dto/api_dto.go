package dto

type ApiSessionDTO struct {
	Id     int    `json:"id,omitempty"`
	Uri    string `json:"uri,omitempty"`
	Method string `json:"method,omitempty"`
}
