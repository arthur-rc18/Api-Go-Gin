package main

import "github.com/gin-gonic/gin"

func ShowAllStudents(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"name": "Lancelot",
	})
}

func main() {

	// The default page of Gin is 8080
	r := gin.Default()
	r.GET("/", ShowAllStudents)
	r.Run()
}
