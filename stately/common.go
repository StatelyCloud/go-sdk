package stately

// ReduceTo is a helper function to unmarshal a list of items of a specific
// item type.
func ReduceTo[TSlice ~[]T, T Item](items ...Item) TSlice {
	// More efficient to not copy by default; only copy if required.
	includeFromI := -1
	for i, v := range items {
		if _, ok := v.(T); !ok {
			includeFromI = i
			break
		}
	}

	// this means nothing failed the predicate and we can return the original slice
	if includeFromI == -1 {
		return nil
	}

	// now create a new slice and copy everything up to the first predicate failure
	result := make(TSlice, includeFromI, len(items))
	for i := 0; i < includeFromI; i++ {
		result[i] = items[i].(T)
	}

	// start iterating from after the first predicate failure and
	// copy everything that passes the predicate
	for i := includeFromI + 1; i < len(items); i++ {
		if _, ok := items[i].(T); ok {
			result = append(result, items[i].(T))
		}
	}

	if len(result) == 0 {
		return nil
	}

	return result
}
