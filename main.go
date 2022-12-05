package main

import (
	"github.com/gin-gonic/gin"
	"golangdojo.com/golangdojo/introtogo/controllers"
	"golangdojo.com/golangdojo/introtogo/initializers"
	// "golangdojo.com/golangdojo/introtogo/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.DELETE("/posts/:id", controllers.PostsDelete)

	r.Run()
}
