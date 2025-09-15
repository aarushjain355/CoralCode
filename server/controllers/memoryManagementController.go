package controllers

import (
	"fmt"
	"net/http"

	"github.com/a9jain/coral-backend/db"
	"github.com/gin-gonic/gin"
)

type SessionID struct {
	ID string `json:"status"`
}

func JoinSession(c *gin.Context) {
	var req SessionID
	if err := c.BindJSON(&req); err != nil {
		fmt.Println("Error parsing JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	fileName, content, err := db.RetrieveFile(req.ID)
	if err != nil {
		fmt.Println("Error retrieving file:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve file"})
		return
	}

	if fileName == "" && content == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	// Return filename and content as JSON
	c.JSON(http.StatusOK, gin.H{
		"filename": fileName,
		"content":  content,
	})
}

func PushToRedis(c *gin.Context) {
	// push contents to redis periodically or by some interval
	fmt.Println("Pushing to Redis...")
}

func Undo(c *gin.Context) {
	// bring last contents from redis with key stroke
	fmt.Println("Pushing to Redis...")
}

func Redo(c *gin.Context) {
	// move a pointer above from redis with a key stroke if possible
	fmt.Println("Pushing to Redis...")
}

func RemoveFromCache(c *gin.Context) {
	// runs an algorithm to remove certain cache items from the redis cache
	fmt.Println("Removing from Cache...")
}
