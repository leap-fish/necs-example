package sliceutil

import "reflect"

func SliceMap[T any, M any](a []T, f func(T) M) []M {
	n := make([]M, len(a))
	for i, e := range a {
		n[i] = f(e)
	}
	return n
}

func TypeList[T any](t []T) []string {
	var strings []string

	for _, entry := range t {
		strings = append(strings, reflect.TypeOf(entry).String())
	}

	return strings
}

// Values returns the values of the map m.
// The values will be in an indeterminate order.
func Values[M ~map[K]V, K comparable, V any](m M) []V {
	r := make([]V, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}
