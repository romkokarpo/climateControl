package main

import (
	"climateControl/server/DTO"
	"climateControl/server/repositories"
	"climateControl/server/services"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	router := gin.Default()

	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/users", handleGetUsers)
	router.GET("/controlSystems", handleGetControlSystems)
	router.GET("/companies", handleGetCompanies)

	router.POST("/updateSysDoc", handleUpdateSysDoc)
	router.POST("/deleteSystemDevices", handleDeleteSystemDevices)
	router.POST("/company/addNewOffices", handleAddNewOfficesToCompany)
	router.POST("/login", handleUserLogin)

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
	var embeddedDevice = DTO.EmbeddedDeviceDto{}
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
	var embeddedDevices []*DTO.EmbeddedDeviceDto

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

func handleGetCompanies(c *gin.Context) {
	repository := repositories.NewCompanyRepository()
	companies := repository.GetAllCompanies()

	c.JSON(http.StatusOK, companies)
}

func handleAddNewOfficesToCompany(c *gin.Context) {
	repository := repositories.NewCompanyRepository()
	decoder := json.NewDecoder(c.Request.Body)
	var newOffices *DTO.EmbeddedOfficeDto

	err := decoder.Decode(&newOffices)
	if err != nil {
		panic(err)
	}

	success := repository.AddNewOffices(newOffices)
	if success {
		c.JSON(http.StatusOK, newOffices)
	} else {
		c.JSON(http.StatusInternalServerError, "Some error occured")
	}
}

func handleUserLogin(c *gin.Context) {
	authService := services.NewAuthenticationService()
	userRepository := repositories.NewUserRepository()
	username := c.PostForm("username")
	password := c.PostForm("password")

	validUserCreds := userRepository.CheckUserCredentials(username, password)
	if !validUserCreds {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User credentials are invalid",
		})
	}

	authService.GenerateToken(username)

	c.SetCookie(
		"id_token",
		authService.JwtToken,
		60*60,
		"/",
		c.Request.Host,
		false,
		true,
	)
	c.JSON(http.StatusOK, bson.M{
		"expiresIn": 60 * 60,
		"idToken":   authService.JwtToken,
	})
}
