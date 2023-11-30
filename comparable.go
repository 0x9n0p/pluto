package pluto

type Comparable interface {
	Comparable() any
}

func Find[T comparable, C Comparable](target T, items ...C) (c C) {
	for _, current := range items {
		if current.Comparable() == target {
			return current
		}
	}
	return
}

func MayFind[T comparable, C Comparable](target T, items ...C) (c C, f bool) {
	for _, current := range items {
		if current.Comparable() == target {
			return current, true
		}
	}
	return
}
