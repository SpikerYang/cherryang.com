package dao

import "cherryang.com/models"

func CreatePost(post models.Post) (err error) {
	result := db.Create(&post)
	return result.Error
}
func FindPosts() (p []models.Post, err error) {
	var posts []models.Post
	result := db.Find(&posts)
	return posts, result.Error
}