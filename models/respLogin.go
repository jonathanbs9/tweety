package models

// ResponseLogin struct => Estructura de respuesta al login
type ResponseLogin struct {
	Token string `json:"token,omitempty"`
}
