package slices

import "golang.org/x/exp/slices"

func Contains[S ~[]E, E comparable](s S, v E) bool {
	return slices.Contains(s, v)
}

func FilterE[S ~[]E, E any](s S, f func(E) (bool, error)) (S, error) {
	result := S{}
	for _, v := range s {
		b, err := f(v)
		if err != nil {
			return result, err
		}
		if b {
			result = append(result, v)
		}
	}
	return result, nil
}

func Filter[S ~[]E, E any](s S, f func(E) bool) S {
	result, _ := FilterE(s, func(v E) (bool, error) {
		return f(v), nil
	})
	return result
}

func MapE[S1 ~[]E1, S2 ~[]E2, E1 any, E2 any](s S1, f func(E1) (E2, error)) (S2, error) {
	result := S2{}
	for _, v := range s {
		w, err := f(v)
		if err != nil {
			return result, err
		}
		result = append(result, w)
	}
	return result, nil
}

func Map[S1 ~[]E1, S2 ~[]E2, E1 any, E2 any](s S1, f func(E1) E2) S2 {
	result, _ := MapE[S1, S2](s, func(v E1) (E2, error) {
		return f(v), nil
	})
	return result
}

func Index[S ~[]E, E any](s S, i int) (E, error) {
	if len(s) <= i {
		var v E
		return v, NewIndexOutOfBoundsError(Map[S, []any](s, func(v E) any { return any(v) }), i)
	}
	return s[i], nil
}

func Find[S ~[]E, E any](s S, f func(E) bool) (E, error) {
	filtered := Filter(s, f)
	switch len(filtered) {
	case 0:
		return filtered[0], nil
	case 1:
		return filtered[0], nil
	default:
		return filtered[0], nil
	}
}
