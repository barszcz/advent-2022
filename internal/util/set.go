package util

type Set[T comparable] struct {
	elemsMap map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	elemsMap := make(map[T]struct{})
	return &Set[T]{
		elemsMap: elemsMap,
	}
}

func NewSetFromSlice[T comparable](elems []T) *Set[T] {
	elemsMap := make(map[T]struct{})
	for _, elem := range elems {
		elemsMap[elem] = struct{}{}
	}
	return &Set[T]{
		elemsMap: elemsMap,
	}
}

func (s *Set[T]) Add(elem T) {
	s.elemsMap[elem] = struct{}{}
}

func (s *Set[T]) Has(elem T) bool {
	_, ok := s.elemsMap[elem]
	return ok
}

func (s *Set[T]) Delete(elem T) {
	delete(s.elemsMap, elem)
}

func (s *Set[T]) Size() int {
	return len(s.elemsMap)
}
