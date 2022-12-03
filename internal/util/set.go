package util

type Set[T comparable] struct {
	elemsMap map[T]struct{}
}

func NewSet[T comparable]() Set[T] {
	elemsMap := make(map[T]struct{})
	return Set[T]{
		elemsMap: elemsMap,
	}
}

func NewSetFromSlice[T comparable](elems []T) Set[T] {
	elemsMap := make(map[T]struct{})
	for _, elem := range elems {
		elemsMap[elem] = struct{}{}
	}
	return Set[T]{
		elemsMap: elemsMap,
	}
}

func (s *Set[T]) Add(elem T) {
	s.elemsMap[elem] = struct{}{}
}

func (s Set[T]) Has(elem T) bool {
	_, ok := s.elemsMap[elem]
	return ok
}

func (s *Set[T]) Delete(elem T) {
	delete(s.elemsMap, elem)
}

func (s Set[T]) Size() int {
	return len(s.elemsMap)
}

func (s Set[T]) Intersection(otherSet Set[T]) Set[T] {
	intersection := NewSet[T]()

	for elem := range s.elemsMap {
		if otherSet.Has(elem) {
			intersection.Add(elem)
		}
	}

	return intersection
}

func (s Set[T]) Union(otherSet Set[T]) Set[T] {
	union := NewSet[T]()

	for elem := range s.elemsMap {
		union.Add(elem)
	}

	for elem := range otherSet.elemsMap {
		union.Add(elem)
	}

	return union
}

func (s Set[T]) Slice() []T {
	slice := make([]T, 0, len(s.elemsMap))

	for elem := range s.elemsMap {
		slice = append(slice, elem)
	}
	return slice
}
