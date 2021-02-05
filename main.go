package main

import (
	"fmt"
	"shortest-path/models"
	"shortest-path/service"
	"shortest-path/storage"
)

func main() {
	storage.SetDefaultStorage(storage.NewInMemoryStorage())

	var path []models.Vertex

	vertices := storage.GetDefaultStorage().GetVertices()
	path = service.FindShortestPathWithMaxEdge(vertices, *vertices["IKA"], *vertices["TLL"], 2)
	if len(path) != 0 {
		panic("Must return empty list")
	}

	path = service.FindShortestPathWithMaxEdge(vertices, *vertices["IKA"], *vertices["TLL"], 3)
	if len(path) != 3 {
		panic("Must return list with three vertices")
	}

	path = service.FindShortestPathWithMaxEdge(vertices, *vertices["IKA"], *vertices["TLL"], 4)
	if len(path) != 3 {
		panic("Must return list with three vertices")
	}

	path = service.FindShortestPath(vertices, *vertices["IKA"], *vertices["TLL"])
	if len(path) != 3 {
		panic("Must return list with three vertices")
	}

	path = service.FindShortestPathWithMaxEdge(vertices, *vertices["IKA"], *vertices["NDB"], 4)
	if len(path) != 0 {
		panic("Must return empty list")
	}

	path = service.FindShortestPathWithMaxEdge(vertices, *vertices["IKA"], *vertices["NDB"], 5)
	if len(path) != 5 {
		panic("Must return list with five vertex")
	}

	path = service.FindShortestPathWithMaxEdge(vertices, *vertices["IKA"], *vertices["NDB"], 6)
	if len(path) != 5 {
		panic("Must return list with five vertex")
	}

	path = service.FindShortestPath(vertices, *vertices["IKA"], *vertices["NDB"])
	if len(path) != 5 {
		panic("Must return list with five vertex")
	}

	fmt.Println("Done")
}
