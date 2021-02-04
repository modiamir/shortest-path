package main

import (
	"fmt"
	"shortest-path/models"
	"shortest-path/parsers"
)

var (
	Vertices map[string]*models.Vertex
)

func main() {
	Vertices = parsers.ParseAirports()
	parsers.ParseAirportRoutes(&Vertices)
	fmt.Println("Done")
}
