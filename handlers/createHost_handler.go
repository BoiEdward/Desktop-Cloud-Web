package handlers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CreateHostPage(c *gin.Context) {
	// Acceder a la sesión
	session := sessions.Default(c)
	usuario := session.Get("usuario").(Persona)

	if usuario.Rol != "Administrador" {
		// Si el usuario no está autenticado, redirige a la página de inicio de sesión
		c.Redirect(http.StatusFound, "/loginPage")
		return
	}

	c.HTML(http.StatusOK, "login.html", nil)
}
