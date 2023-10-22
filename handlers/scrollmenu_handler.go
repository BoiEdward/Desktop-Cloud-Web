package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Scrollmenu(c *gin.Context) {
	c.HTML(http.StatusOK, "scrollmenu.html", gin.H{})
}
