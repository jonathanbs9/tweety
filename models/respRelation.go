package models

// ResponseQueryRelation => Devuelve true/false si existe relación entre 2 usuarios
type ResponseQueryRelation struct {
	Status bool `json:"status"`
}
