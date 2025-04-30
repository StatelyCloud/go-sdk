package stately

// ListResponse allows you to write idiomatic Go code to iterate over a list
// of value. For example, to iterate over a list of items:
//
//	 for listResponse.Next() {
//		 value := listResponse.Value()
//		 // do something with item
//	 }
//	 token, err := listResponse.Token();
//	 // handle error and token
type ListResponse[T any] interface {
	// Next reads an item of the stream, and populates Value() with the current
	// item. It returns true if an item was read and is ready in Value().
	Next() bool
	// Token will only return something when Next() == false, or in other
	// words, the iteration is done. Either a token or an error will be
	// returned. The resulting token can be used for subsequent list calls, or
	// for the sync api.
	Token() (*ListToken, error)
	// Value returns the current item in the iteration.
	Value() T
}
