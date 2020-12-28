package main

import (
	"cherryang.com/controllers"
	"cherryang.com/dao"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	loadHtmlTemplates(router)
	registerHandler(router)
	setStaticFilePath(router)
	initDataSource()
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
	router.GET("/index", controllers.IndexPage)
	blogsRouteGroup := router.Group("/blogs")
	{
		blogsRouteGroup.GET("view", controllers.ViewBlogPage)
		blogsRouteGroup.GET("edit", controllers.EditBlogPage)
		blogsRouteGroup.GET("", controllers.FindBlogs)
		blogsRouteGroup.POST("", controllers.CreateBlog)
	}
	router.GET("/profile", controllers.ProfilePage)
}

func initDataSource() {
	dao.InitDatabase("root:Yqr199733dlj*@tcp(localhost:3306)/blog?charset=utf8&parseTime=True&loc=Local")
}