package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var supEmail string

func MainPage(c *gin.Context) {
	// Acceder a la sesión
	session := sessions.Default(c)
	email := session.Get("email")

	if email == nil {
		// Si el usuario no está autenticado, redirige a la página de inicio de sesión
		c.Redirect(http.StatusFound, "/login")
		return
	}

	supEmail = email.(string)
	machines, _ := consultarMaquinas(email.(string))

	c.HTML(http.StatusOK, "mainPage.html", gin.H{
		"email":    email,
		"machines": machines,
	})
}

func MainSend(c *gin.Context) {
	// Obtener los datos del formulario
	vmname := c.PostForm("vmname")
	os := c.PostForm("os")
	memory, err := strconv.Atoi(c.PostForm("memory"))

	// Crear una estructura Account y convertirla a JSON
	Specifications := Maquina_virtual{Nombre: vmname, Sistema_operativo: os, Memoria: memory, Cpu: 4, Persona_email: supEmail}

	jsonData, err := json.Marshal(Specifications)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if sendJSONMachineToServer(jsonData) {
		// Llamar a consultarMaquinas para obtener la lista actualizada
		maquinas, err := consultarMaquinas(supEmail)
		if err != nil {
			// Manejar el error
		}

		// Redirigir a la página principal con la lista de máquinas actualizada
		c.HTML(http.StatusOK, "mainPage.html", gin.H{
			"machines": maquinas,
		})
	} else {
		// Manejar el caso en el que sendJSONMachineToServer devuelve falso
	}
}

func sendJSONMachineToServer(jsonData []byte) bool {
	serverURL := "http://localhost:8081/json/createVirtualMachine" // Cambia esto por la URL de tu servidor en el puerto 8081

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

func consultarMaquinas(email string) ([]Maquina_virtual, error) {
	serverURL := "http://localhost:8081/json/consultMachine" // Cambia esto por la URL de tu servidor en el puerto 8081

	persona := Persona{Email: email}
	jsonData, err := json.Marshal(persona)
	if err != nil {
		return nil, err
	}

	// Crea una solicitud HTTP POST con el JSON como cuerpo
	req, err := http.NewRequest("POST", serverURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	// Establece el encabezado de tipo de contenido
	req.Header.Set("Content-Type", "application/json")

	// Realiza la solicitud HTTP
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Verifica la respuesta del servidor (resp.StatusCode) aquí si es necesario
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("La solicitud al servidor no fue exitosa")
	}

	// Lee la respuesta del cuerpo de la respuesta HTTP
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var machines []Maquina_virtual

	// Decodifica los datos de respuesta en la variable machines.
	if err := json.Unmarshal(responseBody, &machines); err != nil {
		// Maneja el error de decodificación aquí
	}

	return machines, nil
}

func PowerMachine(c *gin.Context) {
	serverURL := "http://localhost:8081/json/startVM" // Cambia esto por la URL de tu servidor en el puerto 8081

	nombre := c.PostForm("nombreMaquina")
	fmt.Println(nombre)

	payload := map[string]interface{}{
		"tipo_solicitud": "start",
		"nombreVM":       nombre,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return
	}

	// Crea una solicitud HTTP POST con el JSON como cuerpo
	req, err := http.NewRequest("POST", serverURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return
	}

	// Establece el encabezado de tipo de contenido
	req.Header.Set("Content-Type", "application/json")

	// Realiza la solicitud HTTP
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	c.Redirect(http.StatusFound, "/mainPage")
}

func DeleteMachine(c *gin.Context) {
	serverURL := "http://localhost:8081/json/deleteVM" // Cambia esto por la URL de tu servidor en el puerto 8081

	nombre := c.PostForm("nombreMaquina")
	fmt.Println(nombre)

	payload := map[string]interface{}{
		"tipo_solicitud": "delete",
		"nombreVM":       nombre,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return
	}

	// Crea una solicitud HTTP POST con el JSON como cuerpo
	req, err := http.NewRequest("POST", serverURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return
	}

	// Establece el encabezado de tipo de contenido
	req.Header.Set("Content-Type", "application/json")

	// Realiza la solicitud HTTP
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	c.Redirect(http.StatusFound, "/mainPage")
}

func ConfigMachine(c *gin.Context) {
	fmt.Println("hola")
}

func GetMachines(c *gin.Context) {
	// Obtener los datos de las máquinas (puedes utilizar la función `consultarMaquinas`)
	machines, err := consultarMaquinas(supEmail) // Asegúrate de ajustar el email
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Devolver los datos en formato JSON
	c.JSON(http.StatusOK, machines)
}

func Logout(c *gin.Context) {
	// Eliminar la información de la sesión
	session := sessions.Default(c)
	session.Delete("email")
	session.Save()

	// Redirigir al usuario a la página de inicio de sesión u otra página
	c.Redirect(http.StatusFound, "/login")
}
