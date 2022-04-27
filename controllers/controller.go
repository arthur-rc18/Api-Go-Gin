package controllers

import (
	"net/http"

	"github.com/arthur-rc18/Api-Go-Gin/database"
	"github.com/arthur-rc18/Api-Go-Gin/models"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
)

func ShowEveryone(c *gin.Context) {
	var knights []models.Knight
	database.DB.Find(&knights)
	c.JSON(200, knights)
}

func Greeting(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"API": "Presenting " + name,
	})
}

// CreateKnight godoc
// @Summary      Create a new knight
// @Description  Route to create a new knight
// @Tags         knights
// @Accept       json
// @Produce      json
// @Param        knight   body      models.Knight  true  "Knight model"
// @Success      200  {object}	int
// @Failure      400  {object}  httputil.HTTPError
// @Router       /knights [post]
func CreateKnight(c *gin.Context) {

	var knight models.Knight
	if err := c.ShouldBindJSON(&knight); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.ValidateKnight(&knight); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Create(&knight)
	c.JSON(http.StatusOK, knight)

}

// SearchKnightByID godoc
// @Summary      Search a knight by the ID
// @Description  Route to find a knight using the ID
// @Tags         knights
// @Accept       json
// @Produce      json
// @Param        knight   body      models.Knight  true  "Knight model"
// @Success      200  {object}  int
// @Failure      400  {object}  httputil.HTTPError
// @Router       /knights/:id [get]
func SearchKnightByID(c *gin.Context) {
	var knight models.Knight

	id := c.Params.ByName("id")
	database.DB.First(&knight, id)

	if knight.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Knight didn't found"})
		return
	}

	c.JSON(http.StatusOK, knight)
}

func DeleteKnight(c *gin.Context) {
	var knight models.Knight
	id := c.Params.ByName("id")

	database.DB.Delete(&knight, id)
	c.JSON(http.StatusOK, gin.H{"data": "Knight deleted with success"})
}

func EditKnight(c *gin.Context) {
	var knight models.Knight
	id := c.Params.ByName("id")
	database.DB.First(&knight, id)

	if err := c.ShouldBindJSON(&knight); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.ValidateKnight(&knight); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&knight).UpdateColumns(knight)
	c.JSON(http.StatusOK, knight)

}

func SearchByName(c *gin.Context) {
	var knight models.Knight
	name := c.Param("name")
	database.DB.Where(&models.Knight{Name: name}).First(&knight)

	if knight.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Knight didn't found"})
		return
	}

	c.JSON(http.StatusOK, knight)
}

func ShowIndex(c *gin.Context) {
	var knights []models.Knight
	database.DB.Find(&knights)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"mensage": "Welcome",
	})
}

func RouteNotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)

}
