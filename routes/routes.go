package routes

import (
	"github.com/arthur-rc18/Api-Go-Gin/controllers"
	"github.com/gin-gonic/gin"

	// The docs are imported directly from my git repository
	docs "github.com/arthur-rc18/Api-Go-Gin/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRequests() {
	// The default port of Gin is 8080
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/knights", controllers.ShowEveryone)
	r.GET("/:name", controllers.Greeting)
	r.POST("/knights", controllers.CreateKnight)
	r.GET("/knights/:id", controllers.SearchKnightByID)
	r.DELETE("/knights/:id", controllers.DeleteKnight)
	r.PATCH("/knights/:id", controllers.EditKnight)
	r.GET("/knights/name/:name", controllers.SearchByName)
	r.GET("/index", controllers.ShowIndex)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// NoRoute adds handlers for NoRoute. It returns a 404 code by default
	r.NoRoute(controllers.RouteNotFound)
	r.Run()
}
