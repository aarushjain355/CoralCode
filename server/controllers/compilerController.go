package controllers

// pseudocode
// 1. write a service which can choose to swap compilers at any time
// 2. for now, only include python, c++, and coral my own compiler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Compiler int

const (
	Pending Compiler = iota
	Coral
	C_PlusPlus
	Python
)

var controllerType Compiler

type LanguageRequest struct {
	Language string `json:"status"`
}

func ChangeLanguage(c *gin.Context) {
	var req LanguageRequest
	if err := c.BindJSON(&req); err != nil {
		fmt.Println("Error parsing JSON:", err)
		c.String(http.StatusBadRequest, "Invalid JSON")
		return
	}

	switch req.Language {
	case "Python":
		controllerType = Python
	case "C++":
		fmt.Println("Task is in progress")
		controllerType = C_PlusPlus
	default:
		fmt.Println("Unknown status")
		controllerType = Coral
	}
}
