package parsers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"shortest-path/models"
	"shortest-path/utils"
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

	for i := 0; i < len(routes); i++ {
		sourceCode := routes[i]["s"]
		destinationCode := routes[i]["d"]
		if (*verticesMap)[sourceCode] != nil && (*verticesMap)[destinationCode] != nil {
			sourcePoint := (*verticesMap)[sourceCode].Point
			destinationPoint := (*verticesMap)[destinationCode].Point

			distance := utils.GeoDistance(sourcePoint.Latitude, sourcePoint.Longitude, destinationPoint.Latitude, destinationPoint.Longitude)

			(*verticesMap)[sourceCode].AddEdge(models.Edge{To: (*verticesMap)[destinationCode], Distance: distance})
		}
	}
}
