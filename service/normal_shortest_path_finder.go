package service

import (
	"container/heap"
	"github.com/modiamir/shortest-path/models"
	"math"
)

type NormalShortestPathFinder struct {
	verticesMap map[string]*models.Vertex
}

func (f NormalShortestPathFinder) Find(from string, to string) ([]models.Vertex, float64) {
	distances := models.Distances{}
	heap.Init(&distances)

	prev := make(map[string]*models.VertexDistance)
	dist := make(map[string]*models.VertexDistance)
	visited := make(map[string]bool)
	for vertexCode := range f.verticesMap {
		vertexDistance := models.VertexDistance{VertexCode: vertexCode, Distance: 0}
		if from != vertexCode {
			vertexDistance.SetDistance(math.Inf(1))
			prev[vertexCode] = nil
		}
		dist[vertexCode] = &vertexDistance
		visited[vertexCode] = false
		heap.Push(&distances, &vertexDistance)
	}

	for len(distances) > 0 {
		current := heap.Pop(&distances).(*models.VertexDistance)
		visited[current.VertexCode] = true

		if current.VertexCode == to {
			break
		}

		for i := 0; i < len(f.verticesMap[current.VertexCode].Edges); i++ {
			edge := f.verticesMap[current.VertexCode].Edges[i]

			if visited[edge.To.Code] {
				continue
			}

			newDistance := edge.Distance + dist[current.VertexCode].Distance

			if newDistance < dist[edge.To.Code].Distance {
				dist[edge.To.Code].SetDistance(newDistance)
				prev[edge.To.Code] = current
				heap.Fix(&distances, dist[edge.To.Code].Index)
			}
		}
	}

	path := make([]models.Vertex, 0)
	reverseCode := to
	val, ok := prev[reverseCode]
	if ok && val != nil {
		path = prependVertex(path, *f.verticesMap[to])
	}
	for ok && val != nil {
		path = prependVertex(path, *f.verticesMap[val.VertexCode])
		val, ok = prev[val.VertexCode]
	}

	return path, dist[to].Distance
}

func prependVertex(x []models.Vertex, y models.Vertex) []models.Vertex {
	x = append(x, models.Vertex{})
	copy(x[1:], x)
	x[0] = y
	return x
}

func NewNormalShortestPathFinder(verticesMap map[string]*models.Vertex) ShortestPathFinderInterface {
	return NormalShortestPathFinder{verticesMap: verticesMap}
}
