package controllers

import (
	"cherryang.com/dao"
	"cherryang.com/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ViewBlogPage(c *gin.Context) {
	c.HTML(http.StatusOK, "blog.html", gin.H{
	})
}
func EditBlogPage(c *gin.Context) {
	c.HTML(http.StatusOK, "editBlog.html", gin.H{
	})
}

func FindBlogs(c *gin.Context) {
	posts, err := dao.FindPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": posts})
}
func FindBlogById(c *gin.Context) {
	post, err := dao.FindPostById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": post})
}
func CreateBlog(c *gin.Context) {
	var post models.Post
	if err:= c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post.Views = 0
	post.Published = true
	err := dao.CreatePost(post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
