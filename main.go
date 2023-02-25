package main

import (
	"go-backend-gin/controllers"
	"go-backend-gin/initializers"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {

	PORT := os.Getenv("PORT")

	r := gin.Default()

	r.POST("/register", controllers.Register)

	r.Run(":" + PORT) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
