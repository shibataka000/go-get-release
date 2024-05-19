package slices

import "slices"

func IndexFunc[S ~[]E, E any](s S, f func(E) bool) int {
	return slices.IndexFunc(s, f)
}

func Map[S2 ~[]E2, S1 ~[]E1, E2 any, E1 any](s S1, f func(E1) E2) S2 {
	s2, _ := MapE[S2](s, func(e E1) (E2, error) {
		return f(e), nil
	})
	return s2
}

func MapE[S2 ~[]E2, S1 ~[]E1, E2 any, E1 any](s S1, f func(E1) (E2, error)) (S2, error) {
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

func Find[S ~[]E, E any](s S, f func(E) bool) (E, error) {
	i := slices.IndexFunc(s, func(e E) bool {
		return f(e)
	})
	if i == -1 {
		var e E
		return e, &NotFoundError{}
	}
	return s[i], nil
}

type NotFoundError struct{}

func (e *NotFoundError) Error() string {
	return "Not found."
}
