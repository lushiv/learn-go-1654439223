package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func homePage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":    "This is the home page",
		"statusCode": 200,
	})
}

func aboutUsPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":    "This is the about us page",
		"statusCode": 200,
	})
}

func main() {
	r := gin.Default()
	r.GET("/home", homePage)
	r.GET("/about-us", aboutUsPage)
	r.Run()
}
