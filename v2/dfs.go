package v2

func TravelDfs[T comparable](size int, start *Vertex[T]) []Vertex[T] {
	discovered := make(map[T]*Vertex[T], size)
	discovered[start.GetValue()] = start

	stack := make([]*Vertex[T], 0, size)
	stack = append(stack, start)

	result := make([]Vertex[T], 0, size)

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

		result = append(result, *current)
	}

	return result
}
