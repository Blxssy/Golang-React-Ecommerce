package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getUserID(c *gin.Context) (uint, error) {
	id, exists := c.Get("user_id")
	if !exists {
		return 0, errors.New("User ID not found")
	}

	uid, ok := id.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return 0, errors.New("Invalid user ID")
	}

	return uid, nil
}
