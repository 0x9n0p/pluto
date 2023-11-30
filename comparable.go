package pluto

func Find[T Comparable](target T, comparable ...T) (c T) {
	for _, current := range comparable {
		if current.Compare(target) {
			return current
		}
	}
	return
}

func MayFind[T Comparable](target T, comparable ...T) (c T, f bool) {
	for _, current := range comparable {
		if current.Compare(target) {
			return current, true
		}
	}
	return
}

type Comparable interface {
	Compare(Comparable) bool
}
