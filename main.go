package main

import (
	"climateControl/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/users", handleGetUsers)
	router.GET("/controlSystems", handleGetControlSystems)

	router.Run(":3000")
}

func handleGetUsers(c *gin.Context) {
	userRepository := repositories.NewUserRepository()
	c.JSON(http.StatusOK, userRepository.GetAllUsers())
}

func handleGetControlSystems(c *gin.Context) {
	controlSystemRepository := repositories.NewControlSystemRepository()
	c.JSON(http.StatusOK, controlSystemRepository.GetAllControlSystems())
}
