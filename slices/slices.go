package slices

import "slices"

func IndexFunc[S ~[]E, E any](s S, f func(E) bool) int {
	return slices.IndexFunc(s, f)
}

func Map[S1 ~[]E1, S2 ~[]E2, E1 any, E2 any](s S1, f func(E1) E2) S2 {
	s2 := S2{}
	for _, e1 := range s {
		e2 := f(e1)
		s2 = append(s2, e2)
	}
	return s2
}

func MapE[S1 ~[]E1, S2 ~[]E2, E1 any, E2 any](s S1, f func(E1) (E2, error)) (S2, error) {
	s2 := S2{}
	for _, e1 := range s {
		e2, err := f(e1)
		if err != nil {
			return nil, err
		}
		s2 = append(s2, e2)
	}
	return s2, nil
}

func Find[S ~[]E, E any](s S, f func(E) bool) E {
	var e E
	return e
}
