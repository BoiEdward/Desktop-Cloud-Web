package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CreateHostPage(c *gin.Context) {
	// Acceder a la sesión
	session := sessions.Default(c)
	rol := session.Get("rol")

	if rol != "Administrador" {
		// Si el usuario no está autenticado, redirige a la página de inicio de sesión
		c.Redirect(http.StatusFound, "/login")
		return
	}

	c.HTML(http.StatusOK, "createHost.html", nil)
}

func CreateHost(c *gin.Context) {
	// Obtener los datos del formulario
	nombreHost := c.PostForm("nameHost")
	ipHost := c.PostForm("ipHost")

	// Crear un objeto Host con los datos del formulario
	host := Host{
		Nombre: nombreHost,
		Ip:     ipHost,
	}

	// Serializar el objeto host como JSON
	jsonData, err := json.Marshal(host)
	if err != nil {
		// Manejar el error, por ejemplo, responder con un error HTTP
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al serializar el objeto Host"})
		return
	}

	// Definir la URL del servidor
	serverURL := "http://localhost:8081/json/addHost" // Cambia esto por la URL de tu servidor en el puerto 8081

	// Realizar una solicitud HTTP POST
	resp, err := http.Post(serverURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		// Manejar el error, por ejemplo, responder con un error HTTP
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al realizar la solicitud HTTP"})
		return
	}
	defer resp.Body.Close()

	// Verificar el código de estado de la respuesta
	if resp.StatusCode != http.StatusOK {
		// Manejar el error, por ejemplo, responder con un error HTTP
		c.JSON(resp.StatusCode, gin.H{"error": "Error en la respuesta del servidor"})
		return
	}

	// Responder con una confirmación o redirección si es necesario
	c.JSON(http.StatusOK, gin.H{"message": "Host creado con éxito"})
}
