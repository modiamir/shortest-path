package main

import (
	"fmt"
	"shortest-path/models"
	"shortest-path/parsers"
	"shortest-path/service"
)

var (
	Vertices map[string]*models.Vertex
)

func main() {
	Vertices = parsers.ParseAirports()
	parsers.ParseAirportRoutes(&Vertices)

	var path []models.Vertex

	path = service.FindShortestPathWithMaxEdge(Vertices, *Vertices["IKA"], *Vertices["TLL"], 2)
	if len(path) != 0 {
		panic("Must return empty list")
	}

	path = service.FindShortestPathWithMaxEdge(Vertices, *Vertices["IKA"], *Vertices["TLL"], 3)
	if len(path) != 3 {
		panic("Must return list with three vertices")
	}

	path = service.FindShortestPathWithMaxEdge(Vertices, *Vertices["IKA"], *Vertices["TLL"], 4)
	if len(path) != 3 {
		panic("Must return list with three vertices")
	}

	path = service.FindShortestPath(Vertices, *Vertices["IKA"], *Vertices["TLL"])
	if len(path) != 3 {
		panic("Must return list with three vertices")
	}

	path = service.FindShortestPathWithMaxEdge(Vertices, *Vertices["IKA"], *Vertices["NDB"], 4)
	if len(path) != 0 {
		panic("Must return empty list")
	}

	path = service.FindShortestPathWithMaxEdge(Vertices, *Vertices["IKA"], *Vertices["NDB"], 5)
	if len(path) != 5 {
		panic("Must return list with five vertex")
	}

	path = service.FindShortestPathWithMaxEdge(Vertices, *Vertices["IKA"], *Vertices["NDB"], 6)
	if len(path) != 5 {
		panic("Must return list with five vertex")
	}

	path = service.FindShortestPath(Vertices, *Vertices["IKA"], *Vertices["NDB"])
	if len(path) != 5 {
		panic("Must return list with five vertex")
	}

	fmt.Println("Done")
}
