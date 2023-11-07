package main

import (
	"AppWeb/handlers"
	"log"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {

	// Parametro del puerto
	args := os.Args[1]
	port := ":" + args

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
	r.GET("/profile", handlers.ProfilePage)
	r.GET("/welcome", handlers.WelcomePage)
	r.GET("/catalog", handlers.CatalogHandler)
	r.GET("/createHost", handlers.CreateHostPage)

	r.GET("/scrollmenu", handlers.Scrollmenu)
	r.GET("/api/machines", handlers.GetMachines)
	r.GET("/controlMachine", handlers.ControlMachine)
	r.GET("actualizaciones-maquinas", handlers.ActualizacionesMaquinas)

	r.POST("/login", handlers.Login)
	r.POST("/signin", handlers.Signin)
	r.POST("/api/createMachine", handlers.MainSend)
	r.POST("/powerMachine", handlers.PowerMachine)
	r.POST("/deleteMachine", handlers.DeleteMachine)
	r.POST("/configMachine", handlers.ConfigMachine)
	r.POST("/api/loginTemp", handlers.LoginTemp)

	r.POST("/cambiar-contenido", handlers.EnviarContenido)

	// Ruta para cerrar sesión
	r.GET("/logout", handlers.Logout)

	// Iniciar la aplicación
	err := r.Run(port)
	if err != nil {
		log.Fatal(err)
	}
}
