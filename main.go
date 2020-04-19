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
	router.POST("/updateSysDoc", handleUpdateSysDoc)

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

func handleUpdateSysDoc(c *gin.Context) {
	controlSystemRepository := repositories.NewControlSystemRepository()
	systemId := c.PostForm("systemId")
	deviceId := c.PostForm("deviceId")
	newSerialNumber := c.PostForm("newSerialNumber")

	success := controlSystemRepository.UpdateSystemDocument(
		systemId,
		deviceId,
		newSerialNumber,
	)

	if success {
		c.JSON(http.StatusOK, newSerialNumber)
	} else {
		c.JSON(http.StatusInternalServerError, "Some error occured")
	}
}
