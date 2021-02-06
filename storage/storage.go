package storage

import "github.com/modiamir/shortest-path/models"

type VertexStorageInterface interface {
	GetVertices() map[string]*models.Vertex
}

var defaultStorage VertexStorageInterface

func SetDefaultStorage(storage VertexStorageInterface) {
	defaultStorage = storage
}

func GetDefaultStorage() VertexStorageInterface {
	return defaultStorage
}
