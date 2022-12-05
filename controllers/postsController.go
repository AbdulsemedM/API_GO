package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golangdojo.com/golangdojo/introtogo/initializers"
	"golangdojo.com/golangdojo/introtogo/models"
)

func PostsCreate(c *gin.Context) {

	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Post": post,
	})
}

func PostsIndex(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)
	c.JSON(http.StatusOK, gin.H{
		"Post": posts,
	})
}

func PostsShow(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	initializers.DB.Find(&post, id)
	c.JSON(http.StatusOK, gin.H{
		"Post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)
	var post models.Post
	initializers.DB.Find(&post, id)
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})
	c.JSON(200, gin.H{
		"Post": post,
	})
}
func PostsDelete(c *gin.Context) {
	id := c.Param("id")
	initializers.DB.Delete(&models.Post{}, id)
	c.Status(200)
}
