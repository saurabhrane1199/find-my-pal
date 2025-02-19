package routes

import (
	"findmypal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Send friend request
func SendFriendRequest(c *gin.Context) {
	sender, _ := c.Get("username")

	var request struct {
		Receiver string `json:"receiver"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := models.SendFriendRequest(sender.(string), request.Receiver)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send friend request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Friend request sent"})
}

// Accept friend request
func AcceptFriendRequest(c *gin.Context) {
	receiver, _ := c.Get("username")

	var request struct {
		Sender string `json:"sender"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := models.AcceptFriendRequest(request.Sender, receiver.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to accept friend request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Friend request accepted"})
}

// Get list of friends
func GetFriends(c *gin.Context) {
	username, _ := c.Get("username")

	friends, err := models.GetFriends(username.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch friends"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"friends": friends})
}
