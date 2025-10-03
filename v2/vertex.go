package v2

type Vertex[T any] struct {
	value       T
	adjacencies []*Vertex[T]
	weight      []float64
	degreeSize  int
}

func New[T any](value T) Vertex[T] {
	return Vertex[T]{
		value:      value,
		degreeSize: 0,
	}
}

func (v *Vertex[T]) AddAdjacency(neighbor *Vertex[T], weight float64) {
	v.adjacencies = append(v.adjacencies, neighbor)
	v.weight = append(v.weight, weight)
	v.degreeSize += 1
}

func (v *Vertex[T]) GetValue() T {
	return v.value
}

func (v *Vertex[T]) GetDegreeSize() int {
	return v.degreeSize
}

func (v *Vertex[T]) GetEdgeSize() int {
	return v.degreeSize / 2
}

func (v *Vertex[T]) GetAdjacencies() []*Vertex[T] {
	return v.adjacencies
}
