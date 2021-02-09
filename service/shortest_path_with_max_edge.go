package service

import (
	"github.com/modiamir/shortest-path/models"
	"math"
	"strconv"
)

type ShortestPathWithMaxEdgeFinder struct {
	normalFinder ShortestPathFinderInterface
	maxEdge      int
}

func (f ShortestPathWithMaxEdgeFinder) Find(from string, to string) ([]models.Vertex, float64) {
	var shortestPath []models.Vertex
	minimumDistance := math.Inf(1)

	for i := 1; i <= f.maxEdge+1; i++ {
		fromCode := from + ":1"
		toCode := to + ":" + strconv.Itoa(i)
		path, distance := f.normalFinder.Find(fromCode, toCode)
		if distance < minimumDistance {
			minimumDistance = distance
			shortestPath = path
		}
	}

	return shortestPath, minimumDistance
}

func createTempVerticesMap(verticesMap map[string]*models.Vertex, maxNode int) map[string]*models.Vertex {
	tempVerticesMap := make(map[string]*models.Vertex)

	for vertexCode := range verticesMap {
		for i := 1; i <= maxNode; i++ {
			tempVertexCode := vertexCode + ":" + strconv.Itoa(i)
			tempVerticesMap[tempVertexCode] = &models.Vertex{
				Code:    tempVertexCode,
				Name:    verticesMap[vertexCode].Name,
				City:    verticesMap[vertexCode].City,
				Country: verticesMap[vertexCode].Country,
				Point:   verticesMap[vertexCode].Point,
				Edges:   []models.Edge{},
			}
		}
	}

	for vertexCode := range verticesMap {
		for _, edge := range verticesMap[vertexCode].Edges {
			for i := 1; i < maxNode; i++ {
				sourceCode := vertexCode + ":" + strconv.Itoa(i)
				destinationCode := edge.To.Code + ":" + strconv.Itoa(i+1)
				tempVerticesMap[sourceCode].AddEdge(models.Edge{To: tempVerticesMap[destinationCode], Distance: edge.Distance})
			}
		}
	}

	return tempVerticesMap
}

func NewShortestPathWithMaxEdgeFinder(verticesMap map[string]*models.Vertex, maxEdge int) ShortestPathFinderInterface {
	tempVerticesMap := createTempVerticesMap(verticesMap, maxEdge+1)
	return ShortestPathWithMaxEdgeFinder{
		normalFinder: NormalShortestPathFinder{verticesMap: tempVerticesMap},
		maxEdge:      maxEdge,
	}
}
