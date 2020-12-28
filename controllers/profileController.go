package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ProfilePage(c *gin.Context) {
	c.HTML(http.StatusOK, "profile.html", gin.H{
	})
}
