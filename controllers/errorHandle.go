package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleServerError(c *gin.Context, err error, data interface{}) {
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
