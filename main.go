package main

import (
	"climateControl/DTO"
	"climateControl/repositories"
	"encoding/json"
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
	router.POST("/deleteSystemDevices", handleDeleteSystemDevices)

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
	decoder := json.NewDecoder(c.Request.Body)
	var embeddedDevice = DTO.EmbeddedDevice{}
	err := decoder.Decode(&embeddedDevice)
	if err != nil {
		panic(err)
	}

	success := controlSystemRepository.UpdateSystemDevice(embeddedDevice)

	if success {
		c.JSON(http.StatusOK, embeddedDevice.DeviceID)
	} else {
		c.JSON(http.StatusInternalServerError, "Some error occured")
	}
}

func handleDeleteSystemDevices(c *gin.Context) {
	repository := repositories.NewControlSystemRepository()
	decoder := json.NewDecoder(c.Request.Body)
	var embeddedDevices []*DTO.EmbeddedDevice

	err := decoder.Decode(&embeddedDevices)
	if err != nil {
		panic(err)
	}

	deletedNumber := repository.DeleteSystemDevices(embeddedDevices)
	if deletedNumber > 0 {
		c.JSON(http.StatusOK, deletedNumber)
	} else {
		c.JSON(http.StatusInternalServerError, "Some error occured")
	}
}
