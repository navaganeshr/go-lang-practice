package main

import (
	gin "github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// get users route
	router.GET("/users", getUsers)
	router.Run()
}

func getUsers(c *gin.Context) {
  //validate request
  userType := c.Query(key:"type")
  if userType == "" {
	c.JSON(http.StatusBadRequest, gin.H{"error": "type is not specified"})
	return
  }
  
}