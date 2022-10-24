package main

import (
	"fmt"

	"github.com/feacx/crud/initializers"
	"github.com/feacx/crud/models"
)

func init() {
	initializers.LoadEnvVaribles()
	initializers.ConnectDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
	fmt.Println("main")
}
