package undirected_v2

import "errors"

type Graph[T any] struct {
	size       int
	vertices   []*Vertex[T]
	degreeSize int
}

func Build[T any](size int, vertices []*Vertex[T]) (*Graph[T], error) {
	if size != len(vertices) {
		return nil, errors.New("defined size is not equal to vertices size")
	}

	var degreeSize int
	for _, vertex := range vertices {
		degreeSize += vertex.GetDegreeSize()
	}

	return &Graph[T]{
		size:       size,
		vertices:   vertices,
		degreeSize: degreeSize,
	}, nil
}

// get number of vertex
// time-complexity O(1)
// space-complexity O(1)
func (g *Graph[T]) Size() int {
	return g.size
}

func (g *Graph[T]) GetDegreeSize() int {
	return g.degreeSize
}

func (g *Graph[T]) GetEdgeSize() int {
	return g.degreeSize / 2
}
