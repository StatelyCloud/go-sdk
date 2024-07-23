package data

// ListResponse allows you to write idiomatic Go code to iterate over a list of value.
// for example, to iterate over a list of items:
//
//	 for listResponse.Next() {
//		 value := listResponse.Value()
//		 // do something with item
//	 }
//	 token, err := listResponse.Token();
//	 // handle error and token
type ListResponse[T any] interface {
	// Next reads an item of the stream, and populates Value() with the current item.
	Next() bool
	// Token returns the current token OR any error that occurred during iteration.
	Token() (*ListToken, error)
	// Value returns the current item in the iteration.
	Value() T
}
