package domain

type Odontologo struct {
	Id        int    `json:"id"`
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	Matricula string `json:"matricula"`
}
