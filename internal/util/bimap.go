package util

type Bimap[T, U comparable] struct {
	forwardMap map[T]U
	inverseMap map[U]T
}

func NewBimap[T, U comparable]() Bimap[T, U] {
	forwardMap := make(map[T]U)
	inverseMap := make(map[U]T)
	return Bimap[T, U]{
		forwardMap,
		inverseMap,
	}
}

func NewBimapFromMap[T, U comparable](inputMap map[T]U) Bimap[T, U] {
	bimap := NewBimap[T, U]()
	for k, v := range inputMap {
		bimap.Add(k, v)
	}

	return bimap
}

func (b *Bimap[T, U]) Add(forwardKey T, forwardVal U) {
	b.forwardMap[forwardKey] = forwardVal
	b.inverseMap[forwardVal] = forwardKey
}

func (b *Bimap[T, U]) Get(key T) (U, bool) {
	val, ok := b.forwardMap[key]
	return val, ok
}

func (b *Bimap[T, U]) InverseGet(key U) (T, bool) {
	val, ok := b.inverseMap[key]
	return val, ok
}
