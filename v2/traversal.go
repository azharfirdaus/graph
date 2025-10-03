package v2

type Traversal[T comparable] interface {
	Travel(size int, start *Vertex[T]) []Vertex[T]
}
