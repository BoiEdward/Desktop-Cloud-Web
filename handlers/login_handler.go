package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LoginPage(c *gin.Context) {
	// Acceder a la sesión
	session := sessions.Default(c)
	email := session.Get("email")

	if email != nil {
		// Si el usuario no está autenticado, redirige a la página de inicio de sesión
		c.Redirect(http.StatusFound, "/mainPage")
		return
	}

	c.HTML(http.StatusOK, "login.html", nil)
}

func Login(c *gin.Context) {
	// Obtener los datos del formulario
	email := c.PostForm("email")
	password := c.PostForm("password")

	// Crear una estructura Account y convertirla a JSON
	persona := Persona{Email: email, Contrasenia: password}
	jsonData, err := json.Marshal(persona)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if sendJSONToServer(jsonData) {
		session := sessions.Default(c)
		session.Set("email", email) // Almacena el nombre de usuario en la sesión
		session.Save()
		// Redirigir al usuario a la página principal
		c.Redirect(302, "/mainPage")
	} else {
		c.Redirect(http.StatusFound, "/login")
	}
}

func LoginTemp(c *gin.Context) {
	session := sessions.Default(c)
	serverURL := "http://localhost:8081/json/createGuestMachine" // Cambia esto por la URL de tu servidor en el puerto 8081

	// Crea una solicitud HTTP POST vacía (sin cuerpo)
	req, err := http.NewRequest("POST", serverURL, nil)
	if err != nil {
	}

	// Realiza la solicitud
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		// Lee el cuerpo de la respuesta
		responseBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error al leer el cuerpo de la respuesta:", err)
			return
		}

		// Convierte el cuerpo de la respuesta en un mapa
		var data map[string]string
		if err := json.Unmarshal(responseBody, &data); err != nil {
			fmt.Println("Error al decodificar el JSON:", err)
			return
		}

		// Accede a los datos del mapa
		mensaje := data["mensaje"]
		fmt.Println("Mensaje recibido:", mensaje)

		session.Set("email", mensaje) // Almacena el nombre de usuario en la sesión
		session.Save()

		c.Redirect(http.StatusSeeOther, "/controlMachine")
	} else {

	}

}

func sendJSONToServer(jsonData []byte) bool {
	serverURL := "http://localhost:8081/json/login" // Cambia esto por la URL de tu servidor en el puerto 8081

	// Crea una solicitud HTTP POST con el JSON como cuerpo
	req, err := http.NewRequest("POST", serverURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return false
	}

	// Establece el encabezado de tipo de contenido
	req.Header.Set("Content-Type", "application/json")

	// Realiza la solicitud HTTP
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	// Verifica la respuesta del servidor (resp.StatusCode) aquí si es necesario
	if resp.StatusCode != http.StatusOK {
		return false
	} else {
		return true
	}
}
