package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/arthur-rc18/Api-Go-Gin/controllers"
	"github.com/arthur-rc18/Api-Go-Gin/database"
	"github.com/arthur-rc18/Api-Go-Gin/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func RoutesSetup() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func CreateKnightMock() {
	knight := models.Knight{Name: "Knight test", Age: "37", Characteristics: "", Weapon: "Spike"}
	database.DB.Create(knight)
	ID = int(knight.ID)
}

func DeleteKnightMock() {
	var knight models.Knight
	database.DB.Delete(&knight, ID)

}

// In Go to write test, it has to be, by default, with the first word as Test and has to receive a pointer to testing 't *testing.T'
func TestVerifingStatusCodeOfGreetingFunction(t *testing.T) {
	r := RoutesSetup()
	r.GET("/:name", controllers.Greeting)

	req, _ := http.NewRequest("GET", "/knights", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "It should be equal")
	mockRes := `{"API":"Presenting knights"}`
	resBody, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, mockRes, string(resBody))
}

func TestListingAllKnights(t *testing.T) {
	database.DbConnection()
	CreateKnightMock()
	defer DeleteKnightMock()
	r := RoutesSetup()
	r.GET("/knights", controllers.ShowEveryone)
	req, _ := http.NewRequest("GET", "/knights", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestSearchByName(t *testing.T) {
	database.DbConnection()
	CreateKnightMock()
	defer DeleteKnightMock()

	r := RoutesSetup()
	r.GET("/knights/name/:name", controllers.SearchByName)
	req, _ := http.NewRequest("GET", "/knights", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)

}

func TestSearchByID(t *testing.T) {
	database.DbConnection()
	CreateKnightMock()
	defer DeleteKnightMock()
	r := RoutesSetup()
	r.GET("/knights/:id", controllers.SearchKnightByID)
	path := "/knights/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", path, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	var knightMock models.Knight
	json.Unmarshal(res.Body.Bytes(), &knightMock)
	assert.Equal(t, "Knight test", knightMock.Name)

}

func TestDeleteKnight(t *testing.T) {
	database.DbConnection()
	CreateKnightMock()
	r := RoutesSetup()
	r.DELETE("/knights/:id", controllers.DeleteKnight)
	path := "/knights/" + strconv.Itoa(ID)

	req, _ := http.NewRequest("DELETE", path, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestEditKnightHandler(t *testing.T) {
	database.DbConnection()
	CreateKnightMock()
	defer DeleteKnightMock()

	r := RoutesSetup()
	r.PATCH("/knights/:id", controllers.EditKnight)

	knight := models.Knight{Name: "Knight test", Age: "37", Characteristics: "", Weapon: "Spike"}
	jsonValue, _ := json.Marshal(knight)
	path := "/knights/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", path, bytes.NewBuffer(jsonValue))
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	var knightMock models.Knight
	json.Unmarshal(res.Body.Bytes(), &knightMock)
	assert.Equal(t, "37", knightMock.Age)
	assert.Equal(t, "spike", knightMock.Weapon)

}
