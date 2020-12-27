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
	router.Static("/js", "static/js")
}

func registerHandler(router *gin.Engine) {
	router.GET("/index", controllers.IndexHandler)
	blogsRouteGroup := router.Group("/blogs")
	{
		blogsRouteGroup.GET("add", controllers.AddBlogHandler)
		blogsRouteGroup.GET("view", controllers.ViewBlogHandler)
		blogsRouteGroup.GET("edit", controllers.EditBlogHandler)
	}
	router.GET("/profile", controllers.ProfileHandler)
}

func welcomeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
	})
}