package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/a9jain/coral-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var ID int = 0

type SaveFileRequest struct {
	Filename string `json:"filename"`
	Content  string `json:"content"`
}

func TestCode(c *gin.Context) {
	// Read JSON payload into a map
	var payload map[string]interface{}

	if err := c.BindJSON(&payload); err != nil {
		// If parsing fails, print raw body
		fmt.Println("Error parsing JSON:", err)
		c.String(http.StatusBadRequest, "Invalid JSON")
		return
	}

	// Print payload to console
	fmt.Println("Received payload:", payload)

	// run compiler toolchain below

	switch controllers.ControllerType {
	case Python:
		// EXECUTE PYTHON CODE BELOW WITH THE CONTENT
	case C_PlusPlus:
		// do something
		// EXECUTE CPLUSPLUS CODE BELOW WITH THE CONTENT
	default:
		// EXECUTE CORAL CODE BELOW WITH THE CONTENT
		// call the bash script under scripts somehow
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Payload received",
	})
}

func DeleteFile(c *gin.Context) {
	// write business logic to retrieve file information from database for current user and delete
}

func CreateNewFile(c *gin.Context) {
	var payload map[string]interface{}

	fmt.Println("Creating new file...")
	if err := c.BindJSON(&payload); err != nil {
		// If parsing fails, print raw body
		fmt.Println("Error parsing JSON:", err)
		c.String(http.StatusBadRequest, "Invalid JSON")
		return
	}

	var newFile = models.File{
		ID:        uuid.New().String(), // correct conversion from int to string
		FileName:  "new_file.go",       // switch with payload data later
		Content:   "",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	ID += 1
	// somehow return ID as part of the response
	models.CreateFile(&newFile)

}

func SaveFile(c *gin.Context) {

	// need to port this over from redis to postgresql in this api
	var req SaveFileRequest
	if err := c.BindJSON(&req); err != nil {
		fmt.Println("Error parsing JSON:", err)
		c.String(http.StatusBadRequest, "Invalid JSON")
		return
	}

	fmt.Printf("Received filename: %s\n", req.Filename)
	fmt.Printf("Received content: %s\n", req.Content)

	models.UpdateFile(req.Filename, req.Content)

}
