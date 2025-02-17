package routes

import (
	"findmypal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Protected endpoint: Store user location
func PostLocation(c *gin.Context) {
	username, _ := c.Get("username")

	var request struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := models.StoreUserLocation(username.(string), request.Latitude, request.Longitude)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store location"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Location stored successfully"})
}

// Protected endpoint: Get nearby users
func GetNearbyUsers(c *gin.Context) {
	lat, _ := strconv.ParseFloat(c.Query("lat"), 64)
	lon, _ := strconv.ParseFloat(c.Query("lon"), 64)
	radius, _ := strconv.ParseFloat(c.Query("radius"), 64)

	users, err := models.GetNearbyUsers(lat, lon, radius)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch nearby users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"nearby_users": users})
}
