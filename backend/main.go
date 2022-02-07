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
	{ID: 1, Name: "Career Day", Organization: "ISA", Date: "2022"},
	{ID: 2, Name: "2 Day workshop", Organization: "ACM", Date: "2022"},
	{ID: 3, Name: "Big Data Session", Organization: "IEEE", Date: "2022"},
}

func getEvents(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, events)
}

func createNewEvent(c *gin.Context) {
	var newEvent event
	if err := c.BindJSON(&newEvent); err != nil {
		return
	}
	events = append(events, newEvent)
	c.IndentedJSON(http.StatusCreated, newEvent)
}
func main() {
	router := gin.Default()
	router.GET("/events", getEvents)
	router.Run("localhost:8080")
}
