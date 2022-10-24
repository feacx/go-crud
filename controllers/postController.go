package controllers

import (
	"net/http"

	"github.com/feacx/crud/initializers"
	"github.com/feacx/crud/models"
	"github.com/gin-gonic/gin"
)

func PostCreate(ctx *gin.Context) {
	// Get the request body
	var body struct {
		Body  string
		Title string
	}
	ctx.Bind(&body)
	// Create a new post
	post := models.Post{Title: body.Title, Content: body.Body}
	// Save the post to the database

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	// Return the post
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Post created" + result.Name(),
		"data":    post,
	})
}

func GetPosts(ctx *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)
	ctx.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func GetPost(ctx *gin.Context) {
	var post models.Post
	id := ctx.Param("id")
	initializers.DB.First(&post, id)
	ctx.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func UpdatePost(ctx *gin.Context) {
	// get id from url
	id := ctx.Param("id")

	// get the data from the request body
	var body struct {
		Body  string
		Title string
	}
	ctx.Bind(&body)

	// find the post
	var post models.Post
	res := initializers.DB.First(&post, id).Error
	if res != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "文章不存在",
		})
		return
	}

	// update the post
	initializers.DB.Model(&post).Updates(models.Post{
		Title:   body.Title,
		Content: body.Body,
	})

	// return the post
	ctx.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func DelPost(ctx *gin.Context) {
	var post models.Post
	id := ctx.Param("id")
	initializers.DB.Delete(&post, id)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Post deleted",
	})
}
