package parsers

import (
	"encoding/json"
	"github.com/modiamir/shortest-path/models"
	"github.com/modiamir/shortest-path/utils"
	"io/ioutil"
	"log"
)

func ParseAirports() map[string]*models.Vertex {
	var vertices []models.Vertex
	fileName := "./dataset/airports.json"
	jsonFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = json.Unmarshal(jsonFile, &vertices)

	verticesMap := make(map[string]*models.Vertex)

	for i := 0; i < len(vertices); i++ {
		vertices[i].Edges = []models.Edge{}
		verticesMap[vertices[i].Code] = &vertices[i]
	}

	return verticesMap
}

func ParseAirportRoutes(verticesMap *map[string]*models.Vertex) {
	fileName := "./dataset/airportRoutes.json"
	jsonFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err.Error())
	}
	var routes []map[string]string

	err = json.Unmarshal(jsonFile, &routes)

	visitedRoutes := make(map[string]bool)

	for i := 0; i < len(routes); i++ {
		sourceCode := routes[i]["s"]
		destinationCode := routes[i]["d"]
		visitedRoute := sourceCode + destinationCode
		if !visitedRoutes[visitedRoute] && (*verticesMap)[sourceCode] != nil && (*verticesMap)[destinationCode] != nil {
			visitedRoutes[visitedRoute] = true
			sourcePoint := (*verticesMap)[sourceCode].Point
			destinationPoint := (*verticesMap)[destinationCode].Point

			distance := utils.GeoDistance(sourcePoint.Latitude, sourcePoint.Longitude, destinationPoint.Latitude, destinationPoint.Longitude)

			(*verticesMap)[sourceCode].AddEdge(models.Edge{To: (*verticesMap)[destinationCode], Distance: distance})
		}
	}
}
