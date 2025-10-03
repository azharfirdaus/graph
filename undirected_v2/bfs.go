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
		neighbors := current.GetNeighbors()
		for _, neighbor := range neighbors {
			_, ok := discovered[neighbor.GetValue()]
			if !ok {
				discovered[neighbor.GetValue()] = neighbor
				queue = append(queue, neighbor)
			}
		}

		queue = queue[1:]
		result = append(result, *current)
	}

	return result
}
