package main

import (
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
)

type event struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Organization string `json:"organization"`
	Date         string `json:"date"`
}

var events = []event{
	{ID: "1", Name: "Career Day", Organization: "ISA", Date: "2022"},
	{ID: "2", Name: "2 Day workshop", Organization: "ACM", Date: "2022"},
	{ID: "3", Name: "Big Data Session", Organization: "IEEE", Date: "2022"},
}

func getEvents(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, events)
}

func eventById(c *gin.Context) {
	id := c.Param("id")
	event, err := getEventById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, event)

}

func getEventById(id string) (*event, error) {
	for index, event := range events {
		if event.ID == id {
			return &events[index], nil
		}

	}
	return nil, errors.New("event not found")
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
	router.POST("/events/new", createNewEvent)
	router.GET("/events/:id", eventById)
	router.Run("localhost:8080")
}
