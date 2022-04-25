package routes

import (
	"github.com/arthur-rc18/Api-Go-Gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	// The default port of Gin is 8080
	r := gin.Default()
	r.GET("/knights", controllers.ShowEveryone)
	r.GET("/:name", controllers.Greeting)
	r.POST("/knights", controllers.CreateKnight)
	r.GET("/knights/:id", controllers.SearchKnightByID)
	r.DELETE("/knights/:id", controllers.DeleteKnight)
	r.Run()
}
