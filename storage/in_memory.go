package storage

import (
	"shortest-path/models"
	"shortest-path/parsers"
)

type InMemoryVertexStorage struct {
	initialized bool
	vertices    map[string]*models.Vertex
}

func (s InMemoryVertexStorage) GetVertices() map[string]*models.Vertex {
	if s.initialized {
		return s.vertices
	}

	s.initialize()

	return s.vertices
}

func (s *InMemoryVertexStorage) initialize() {
	s.vertices = parsers.ParseAirports()
	parsers.ParseAirportRoutes(&s.vertices)
	s.initialized = true
}

func NewInMemoryStorage() InMemoryVertexStorage {
	storage := InMemoryVertexStorage{}
	storage.initialize()

	return storage
}
