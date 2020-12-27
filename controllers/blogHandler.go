package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ViewBlogHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "blog.html", gin.H{
	})
}
func EditBlogHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "editBlog.html", gin.H{
	})
}
func AddBlogHandler(c *gin.Context) {

}
