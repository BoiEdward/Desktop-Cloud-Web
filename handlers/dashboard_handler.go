package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DatosCatalogo representa los datos para el catálogo de máquinas virtuales
type DatosCatalogo struct {
	TotalMaquinas       int // Total de máquinas virtuales
	TotalEncendidas     int
	UsuariosRegistrados int // Total de usuarios registrados
	TotalCPU            int // Total de CPU utilizado
	TotalRAM            int // Total de RAM utilizada
}

func DashboardHandler(c *gin.Context) {
	// ... tu código existente ...

	// Calcula los datos para el catálogo (esto es solo un ejemplo, debes obtener estos datos de tu lógica)
	datosCatalogo := DatosCatalogo{
		TotalMaquinas:       100,
		TotalEncendidas:     25,
		UsuariosRegistrados: 50,
		TotalCPU:            75,
		TotalRAM:            200,
	}

	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"email":          "email",
		"machines":       nil,
		"machinesChange": nil,
		"datosCatalogo":  datosCatalogo,
	})
}
