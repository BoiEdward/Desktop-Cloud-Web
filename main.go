package main

import (
	"AppWeb/handlers"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Configurar la tienda de cookies para las sesiones
	store := cookie.NewStore([]byte("tu_clave_secreta"))
	r.Use(sessions.Sessions("sesion", store))

	// Configura las rutas
	r.LoadHTMLGlob("templates/*.html")
	r.Static("/static", "./static")

	r.GET("/login", handlers.LoginPage)
	r.GET("/signin", handlers.SigninPage)
	r.GET("/mainPage", handlers.MainPage)
	r.GET("/navbar", handlers.NavbarPage)
	r.GET("/scrollmenu", handlers.Scrollmenu)
	r.GET("/api/machines", handlers.GetMachines)

	r.POST("/login", handlers.Login)
	r.POST("/signin", handlers.Signin)
	r.POST("/mainPage", handlers.MainSend)
	r.POST("/powerMachine", handlers.PowerMachine)
	r.POST("/deleteMachine", handlers.DeleteMachine)
	r.POST("/configMachine", handlers.ConfigMachine)

	// Ruta para cerrar sesión
	r.GET("/logout", handlers.Logout)

	// Ruta para cerrar sesión
	//r.GET("/logout", handlers.Logout)

	// Iniciar la aplicación
	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
