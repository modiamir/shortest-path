package models

type VertexDistance struct {
	Distance   float64
	VertexCode string
	Index      int
}

func (vd *VertexDistance) SetDistance(distance float64) {
	vd.Distance = distance
}

func (vd *VertexDistance) SetIndex(index int) {
	vd.Index = index
}

type Distances []*VertexDistance

func (d Distances) Len() int {
	return len(d)
}

func (d Distances) Less(i, j int) bool {
	return d[i].Distance < d[j].Distance
}

func (d *Distances) Swap(i, j int) {
	(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
	(*d)[i].SetIndex(i)
	(*d)[j].SetIndex(j)
}

func (d *Distances) Push(x interface{}) {
	x.(*VertexDistance).SetIndex(len(*d))
	*d = append(*d, x.(*VertexDistance))
}

func (d *Distances) Pop() interface{} {
	old := *d
	n := len(old)
	x := old[n-1]
	*d = old[0 : n-1]
	return x
}
