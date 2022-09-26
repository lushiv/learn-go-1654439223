package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//1. Create our data structure and mock data
type Users struct {
	ID       int64  `json: "id"`
	UUID     string `json: "uuid"`
	NAME     string `json: "name"`
	PASSWORD string `json: "password"`
	PHONE    string `json: "phone"`
	STATUS   bool   `json: "status"`
}

// making array
var usersList = []Users{
	{ID: 1, UUID: "d24d4e54-3d3e-11ed-be6e-67909d485b89", NAME: "Janak", PASSWORD: "Test@1234", PHONE: "9865054974", STATUS: true},
	{ID: 2, UUID: "d685dfcc-3d3e-11ed-82b9-1f4e5fb4388d", NAME: "Jhon", PASSWORD: "Test@1234", PHONE: "9865054974", STATUS: true},
	{ID: 3, UUID: "da2c6204-3d3e-11ed-97ee-0b4d252c078b", NAME: "Ronit", PASSWORD: "Test@1234", PHONE: "9865054974", STATUS: true},
	{ID: 4, UUID: "dffb2080-3d3e-11ed-91a5-ab33c3e34fb2", NAME: "Racky", PASSWORD: "Test@1234", PHONE: "9865054974", STATUS: true},
}

func getUsersList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":    "operation successfull",
		"statusCode": 200,
		"data":       usersList,
	})
}

func createUser(c *gin.Context) {
	var input_users_data Users

	if err := c.ShouldBindJSON(&input_users_data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	usersList = append(usersList, input_users_data)
	c.JSON(http.StatusOK, gin.H{
		"message":    "operation successfull",
		"statusCode": 200,
		"data":       usersList,
	})

}

func deleteUser(c *gin.Context) {
	uuid := c.Param("uuid")
	println(uuid)

	for index, element := range usersList {
		println(index)
		if uuid == element.UUID {
			c.JSON(http.StatusOK, gin.H{
				"message":    "delete successfull",
				"statusCode": 200,
			})

		}
		c.JSON(http.StatusNoContent, gin.H{
			"message": "operation failed",
		})

	}

}

func main() {
	router := gin.Default()

	router.GET("/users", getUsersList)
	router.POST("/create", createUser)
	router.DELETE("/delete/:uuid", deleteUser)

	router.Run()
}
