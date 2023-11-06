package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CatalogItem struct {
	ID                           int
	Nombre                       string
	SistemaOperativo             string
	DistribucionSistemaOperativo string
	Ram                          int
	Cpu                          int
}

func CatalogHandler(c *gin.Context) {
	// Supongamos que tienes una lista de ítems en tu catálogo
	catalog := []CatalogItem{
		{1, "Máquina 1", "Linux", "Ubuntu", 4, 2},
		{2, "Máquina 2", "Windows", "10", 8, 4},
		// ... otros elementos del catálogo
	}

	c.HTML(http.StatusOK, "catalog.html", gin.H{
		"Catalog": catalog,
	})
}

func CreateVMHandler(c *gin.Context) {
	// Obtiene el ID de la máquina virtual seleccionada desde la URL
	//id := c.Param("id")
	// Aquí puedes realizar las acciones necesarias para crear la máquina virtual con el ID seleccionado
	// ...

	// Ejemplo de respuesta después de crear la máquina virtual
	c.JSON(http.StatusOK, gin.H{"mensaje": "Máquina virtual creada exitosamente"})
}
