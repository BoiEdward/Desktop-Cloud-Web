package handlers

type Account struct {
	Username string `json:"nombre"`
	Email    string `json:"correo"`
	Password string `json:"contrasenia"`
}

type Persona struct {
	Nombre      string
	Apellido    string
	Email       string
	Contrasenia string
}

type Specifications struct {
	Name   string `json:"nombre"`
	OSType string `json:"tipoSO"`
	Memory int    `json:"memoria"`
	CPU    int    `json:"cpu"`
}

type Maquina_virtual struct {
	Uuid              string
	Nombre            string
	Sistema_operativo string
	Memoria           int
	Cpu               int
	Estado            string
	Hostname          string
	Ip                string
	Persona_email     string
}
