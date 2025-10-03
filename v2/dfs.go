package v2

type StackDFSTraversal[T comparable] struct {
}

// travel from given vertex using Depth-First Search
//
// time-complexity O(V + E)
//
//	where V is the number of vertices
//	where E is the number of edges
//
// space-complexity O(V)
//
//	where V is the number of vertices
func (s *StackDFSTraversal[T]) Travel(size int, start *Vertex[T]) []Vertex[T] {
	discovered := make(map[T]*Vertex[T], size)
	discovered[start.GetValue()] = start

	stack := make([]*Vertex[T], 0, size)
	stack = append(stack, start)

	path := make([]Vertex[T], 0, size)

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		adjacencies := current.GetAdjacencies()
		for _, adjacency := range adjacencies {
			_, ok := discovered[adjacency.GetValue()]
			if !ok {
				discovered[adjacency.GetValue()] = adjacency
				stack = append(stack, adjacency)
			}
		}

		path = append(path, *current)
	}

	return path
}
