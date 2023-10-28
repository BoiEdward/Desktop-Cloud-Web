package handlers

type Persona struct {
	Nombre      string
	Apellido    string
	Email       string
	Contrasenia string
	Rol         string
}

type Maquina_virtual struct {
	Uuid                           string
	Nombre                         string
	Ram                            int
	Cpu                            int
	Ip                             string
	Estado                         string
	Hostname                       string
	Persona_email                  string
	Host_id                        int
	Disco_id                       int
	Sistema_operativo              string
	Distribucion_sistema_operativo string
}
