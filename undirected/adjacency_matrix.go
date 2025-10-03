package undirected

import "strings"

type AdjacencyMatrix [][]bool

func New(size int) AdjacencyMatrix {
	adjacencyMatrix := make([][]bool, size)
	for index, _ := range adjacencyMatrix {
		adjacencyMatrix[index] = make([]bool, size)
	}
	return adjacencyMatrix
}

func (a AdjacencyMatrix) String() string {
	var sb strings.Builder
	for _, row := range a {
		for _, isEdgeEstablished := range row {
			if isEdgeEstablished {
				sb.WriteString("1")
			} else {
				sb.WriteString("0")
			}
			sb.WriteByte(byte(' '))
		}
		sb.WriteByte(byte('\n'))
	}
	return sb.String()
}
