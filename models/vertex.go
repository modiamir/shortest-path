package models

type Vertex struct {
	Code    string `json:"code"`
	Name    string `json:"name"`
	City    string `json:"city"`
	Country string `json:"country"`
	Point   Point  `json:"_geoloc"`
	Edges   []Edge `json:"-"`
}

type VertexList struct {
	Vertices []Vertex
}

func (v *Vertex) AddEdge(edge Edge) {
	if v.Edges == nil {
		v.Edges = []Edge{}
	}

	v.Edges = append(v.Edges, edge)
}
