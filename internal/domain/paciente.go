package domain

type Paciente struct {
	Id        int    `json:"id"`
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	Domicilio string `json:"domicilio"`
	DNI       int    `json:"dni"`
	FechaAlta string `json:"fechaAlta"`
}
