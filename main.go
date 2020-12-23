package main

import (
	"cherryang.com/controllers"
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
	router.Static("/templates", "static/templates")
	router.Static("/css", "static/css")
	router.Static("/images", "static/images")
}

func registerHandler(router *gin.Engine) {
	router.GET("/", controllers.IndexHandler)
}

func welcomeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
	})
}