package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Footballer struct {
	ID          string `JSON:"id"`
	Name        string `JSON:"name"`
	ShirtNumber uint16 `JSON:"number"`
	Age         uint16 `JSON:"age"`
	Nationality string `JSON:"nationality"`
}

var madridteam = []Footballer{
	{ID: "1", Name: "Thibaut Courtois", ShirtNumber: 1, Age: 31, Nationality: "Belgium"},
	{ID: "2", Name: "Eder Militao", ShirtNumber: 3, Age: 25, Nationality: "Brazil"},
	{ID: "3", Name: "Antonio Rudiger", ShirtNumber: 22, Age: 30, Nationality: "Germany"},
	{ID: "4", Name: "Nacho Fernandez", ShirtNumber: 6, Age: 33, Nationality: "Spain"},
	{ID: "5", Name: "Jesus Vallejo", ShirtNumber: 24, Age: 26, Nationality: "Spain"},
	{ID: "6", Name: "Ferland Mendy", ShirtNumber: 23, Age: 28, Nationality: "France"},
	{ID: "7", Name: "Fran Garcia", ShirtNumber: 1, Age: 23, Nationality: "Spain"}, // Assuming - represents empty shirt number
	{ID: "8", Name: "Daniel Carvajal", ShirtNumber: 2, Age: 31, Nationality: "Spain"},
	{ID: "9", Name: "Lucas Vazquez", ShirtNumber: 17, Age: 32, Nationality: "Spain"},
	{ID: "10", Name: "Alvaro Odriozola", ShirtNumber: 16, Age: 27, Nationality: "Spain"},
}

// getAlbums responds with the list of all albums as JSON.
func getPlayers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, madridteam)
}
func get_player_by_id(c *gin.Context) {
	id := c.Param("id")

	for _, value := range madridteam {
		if value.ID == id {
			c.IndentedJSON(http.StatusOK, value)
			return
		}
	}
	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "can not find the player"})
}
func post_a_player(c *gin.Context) {
	var new_player Footballer

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&new_player); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "can not add a player"})
		return
	}

	//add a new player to slice
	madridteam = append(madridteam, new_player)
	c.IndentedJSON(http.StatusOK, new_player)
}

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/players/", getPlayers)
		v1.POST("/players/", post_a_player)
		v1.GET("/players/:id", get_player_by_id)
	}
	router.Run("localhost:8080")
}
