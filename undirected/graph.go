package undirected

import (
	"errors"

	"github.com/golang-ds/linkedlist/singly"
)

type Vertex[T any] struct {
	value T
}

type Edge[T any] struct {
	Begin *Vertex[T]
	End   *Vertex[T]
}

type Graph[T any] struct {
	size     int
	vertices []*Vertex[T]
	edges    []*Edge[T]
}

// build and returns an undirected immutable graph
// that contain vertices and edges of T respectively
// vertices should not have duplication
// time complexity O(1)
// space complexity O(v + e)
func Build[T any](size int, vertices []*Vertex[T], edges ...*Edge[T]) (*Graph[T], error) {
	if size != len(vertices) {
		return nil, errors.New("defined size is not equal to vertices size")
	}

	return &Graph[T]{
		size:     size,
		vertices: vertices,
		edges:    edges,
	}, nil
}

// get number of vertex
// time-complexity O(1)
// space-complexity O(1)
func (g *Graph[T]) Size() int {
	return g.size
}

// get the copy of adjacencyMatrix
// time-complexity O(v + e), O(quadratic v) for dense graphs.
// space-complexity O(quadratic v)
func (g *Graph[T]) GetAdjacencyMatrix() AdjacencyMatrix {
	matrix := New(g.size)

	vertexIndex := make(map[*Vertex[T]]int, g.size)

	for index, vertex := range g.vertices {
		vertexIndex[vertex] = index
	}

	for _, edge := range g.edges {
		beginIndex, beginFound := vertexIndex[edge.Begin]
		endIndex, endFound := vertexIndex[edge.End]

		if beginFound && endFound {
			matrix[beginIndex][endIndex] = true
		}
	}

	return matrix
}

func (g *Graph[T]) GetAdjacencyList() []singly.LinkedList[*Vertex[T]] {
	slice := make([]singly.LinkedList[*Vertex[T]], g.size)
	for index, vertex := range g.vertices {
		for _, edge := range g.edges {
			list := singly.New[*Vertex[T]]()
			if vertex == edge.Begin {
				list.AddFirst(edge.End)
			}
			slice[index] = list
		}
	}

	return slice
}
