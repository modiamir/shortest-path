package service

import (
	"container/heap"
	"fmt"
	"math"
	"shortest-path/models"
)

func FindShortestPath(verticesMap map[string]*models.Vertex, from models.Vertex, to models.Vertex) []models.Vertex {
	distances := models.Distances{}
	heap.Init(&distances)

	prev := make(map[string]*models.VertexDistance)
	dist := make(map[string]*models.VertexDistance)
	for vertexCode := range verticesMap {
		vertexDistance := models.VertexDistance{VertexCode: vertexCode, Distance: 0}
		if from.Code != vertexCode {
			vertexDistance.SetDistance(math.Inf(1))
			prev[vertexCode] = nil
		}
		dist[vertexCode] = &vertexDistance
		heap.Push(&distances, &vertexDistance)
	}

	for len(distances) > 0 {
		current := heap.Pop(&distances).(*models.VertexDistance) // u

		if current.VertexCode == to.Code {
			break
		}

		for i := 0; i < len(verticesMap[current.VertexCode].Edges); i++ {
			edge := verticesMap[current.VertexCode].Edges[i] // v's edge
			newDistance := edge.Distance + dist[current.VertexCode].Distance
			if newDistance < dist[edge.To.Code].Distance {
				dist[edge.To.Code].SetDistance(newDistance)
				prev[edge.To.Code] = current
				heap.Fix(&distances, dist[edge.To.Code].Index)
			}
		}
		vc := current.VertexCode
		fmt.Println(vc)
	}

	path := make([]models.Vertex, 0)
	reverseCode := to.Code
	val, ok := prev[reverseCode]
	if ok {
		path = prependVertex(path, *verticesMap[to.Code])
	}
	for ok {
		path = prependVertex(path, *verticesMap[val.VertexCode])
		val, ok = prev[val.VertexCode]
	}

	return path
}

func prependVertex(x []models.Vertex, y models.Vertex) []models.Vertex {
	x = append(x, models.Vertex{})
	copy(x[1:], x)
	x[0] = y
	return x
}
