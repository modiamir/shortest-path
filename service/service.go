package service

import "github.com/modiamir/shortest-path/models"

type ShortestPathFinderInterface interface {
	Find(from string, to string) ([]models.Vertex, float64)
}
