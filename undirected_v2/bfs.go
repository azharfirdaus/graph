package undirected_v2

// build graph by given size and array of vertices
//
// time-complexity O(V + E)
//
//	where V is the number of vertices
//	where E is the number of edges
//
// space-complexity O(V)
//
//	where V is the number of vertices
func TravelBfs[T comparable](size int, start *Vertex[T]) []Vertex[T] {
	discovered := make(map[T]*Vertex[T], size)
	discovered[start.GetValue()] = start

	queue := make([]*Vertex[T], 0, size)
	queue = append(queue, start)

	result := make([]Vertex[T], 0, size)

	for len(queue) > 0 {
		current := queue[0]
		adjacencies := current.GetAdjacencies()
		for _, adjacency := range adjacencies {
			_, ok := discovered[adjacency.GetValue()]
			if !ok {
				discovered[adjacency.GetValue()] = adjacency
				queue = append(queue, adjacency)
			}
		}

		queue = queue[1:]
		result = append(result, *current)
	}

	return result
}
