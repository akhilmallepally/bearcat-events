package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type event struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Organization string `json:"organization"`
	Date         string `json:"date"`
}

var events = []event{
	{ID: 1, Name: "Shutter Island", Organization: "Martin Scorsese", Date: "2010"},
	{ID: 2, Name: "Shutter Island", Organization: "Martin Scorsese", Date: "2010"},
	{ID: 3, Name: "Shutter Island", Organization: "Martin Scorsese", Date: "2010"},
}

func getEvents(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, events)
}

func main() {
	router := gin.Default()
	router.GET("/movies", getEvents)
	router.Run("localhost:8080")
}
