package main

import (
	"github.com/feacx/crud/controllers"
	"github.com/feacx/crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVaribles()
	initializers.ConnectDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostCreate)
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetPost)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DelPost)
	r.Run()
}
