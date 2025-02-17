package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProtectedRoute(c *gin.Context) {
	username, _ := c.Get("username")
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the protected route", "user": username})
}
