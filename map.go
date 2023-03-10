package slice

// Keys returns a slice of keys extracted from the input map m. The keys are returned
// in no particular order. It returns nil if the map is empty.
func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	if len(m) == 0 {
		return nil
	}

	s := make([]K, 0, len(m))

	for k := range m {
		s = append(s, k)
	}

	return s
}

// Values returns a slice of values extracted from the input map m. The values are returned
// in no particular order. It returns nil if the map is empty.
func Values[M ~map[K]V, K comparable, V any](m M) []V {
	if len(m) == 0 {
		return nil
	}

	s := make([]V, 0, len(m))

	for _, v := range m {
		s = append(s, v)
	}

	return s
}
