package service

import (
	"container/heap"
	"math"
	"shortest-path/models"
)

func FindShortestPathWithMaxEdge(
	verticesMap map[string]*models.Vertex,
	from models.Vertex,
	to models.Vertex,
	maxEdge int) []models.Vertex {
	return dijkstra(verticesMap, from, to, maxEdge)
}

func FindShortestPath(
	verticesMap map[string]*models.Vertex,
	from models.Vertex,
	to models.Vertex) []models.Vertex {
	return dijkstra(verticesMap, from, to, 0)
}

func dijkstra(verticesMap map[string]*models.Vertex, from models.Vertex, to models.Vertex, maxEdge int) []models.Vertex {
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
		current := heap.Pop(&distances).(*models.VertexDistance)

		if current.VertexCode == to.Code {
			break
		}

		for i := 0; i < len(verticesMap[current.VertexCode].Edges); i++ {
			edge := verticesMap[current.VertexCode].Edges[i]
			newDistance := edge.Distance + dist[current.VertexCode].Distance
			if newDistance < dist[edge.To.Code].Distance && (maxEdge <= 0 || (dist[current.VertexCode].FlightCount+1) < maxEdge) {
				if maxEdge > 0 {
					dist[edge.To.Code].SetFlightCount(dist[current.VertexCode].FlightCount + 1)
				}
				dist[edge.To.Code].SetDistance(newDistance)
				prev[edge.To.Code] = current
				heap.Fix(&distances, dist[edge.To.Code].Index)
			}
		}
	}

	path := make([]models.Vertex, 0)
	reverseCode := to.Code
	val, ok := prev[reverseCode]
	if ok && val != nil {
		path = prependVertex(path, *verticesMap[to.Code])
	}
	for ok && val != nil {
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
