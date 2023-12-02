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

// MayFind returns -1 if the item not found.
func MayFind[T comparable, C Comparable](target T, items ...C) (c C, i int) {
	for i, current := range items {
		if current.Comparable() == target {
			return current, i
		}
	}
	return c, -1
}
