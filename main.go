package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	"go-app/controllers"
	"go-app/inits"
	"go-app/middlewares"
)

func init() {
	inits.LoadEnv()
	inits.DBInit()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	r.POST("/user", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/auth", middlewares.RequireAuth, controllers.GetAuthenticatedUser)
	r.GET("/users", middlewares.RequireAuth, controllers.GetUsers)

	fmt.Printf("Server is now running at http://localhost:%s", os.Getenv("PORT"))

	r.Run(":8000")
}
