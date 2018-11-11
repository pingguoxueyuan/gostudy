package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginViewHandle(c *gin.Context) {
	c.HTML(http.StatusOK, "views/login.html", nil)
}

func RegisterViewHandle(c *gin.Context) {
	c.HTML(http.StatusOK, "views/register.html", nil)
}
