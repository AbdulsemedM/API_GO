package main

import (
	"golangdojo.com/golangdojo/introtogo/initializers"
	"golangdojo.com/golangdojo/introtogo/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})

}
