package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	loadHtmlTemplates(router)
	registerHandler(router)
	setStaticFilePath(router)
	router.Run()
}

func loadHtmlTemplates(router *gin.Engine) {
	router.LoadHTMLGlob("static/templates/*")
}

func setStaticFilePath(router *gin.Engine) {
	router.Static("/css", "static/css")
	router.Static("/images", "static/images")
}

func registerHandler(router *gin.Engine) {
	router.GET("/welcome", welcomeHandler)
}

func welcomeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
	})
}