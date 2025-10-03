package undirected_v2

type Vertex[T any] struct {
	value      T
	neighbors  []*Vertex[T]
	degreeSize int
}

func New[T any](value T) Vertex[T] {
	return Vertex[T]{
		value:      value,
		degreeSize: 0,
	}
}

func (v *Vertex[T]) AddNeighbor(neighbor *Vertex[T]) {
	v.neighbors = append(v.neighbors, neighbor)
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
