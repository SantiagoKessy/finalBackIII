package domain

type Turno struct {
	Id          int        `json:"id"`
	Odontologo  Odontologo `json:"odontologo"`
	Paciente    Paciente   `json:"paciente"`
	Fecha       string     `json:"fecha"`
	Hora        string     `json:"hora"`
	Descripcion string     `json:"descripcion"`
}
