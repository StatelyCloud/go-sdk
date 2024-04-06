package data

// IsTransactionCommand is a type that exports a private type that proto people refuse
// to export. Exporting this makes life significantly easier.
// https://github.com/golang/protobuf/pull/588
// https://github.com/golang/protobuf/issues/261#issuecomment-430496210
type IsTransactionCommand = isTransactionRequest_Command

// IsTransactionResponseResult is a type that exports a private type that proto people refuse
// to export. Exporting this makes life significantly easier.
type IsTransactionResponseResult = isTransactionResponse_Result

// IsListResponder is a type that binds ListPartialResult to the ListResponder interface.
func (*ListPartialResult) IsListResponder() {}

// IsListResponder is a type that binds ListFinished to the ListResponder interface.
func (*ListFinished) IsListResponder() {}

// ListResponder allows us to conform both txn and non-txn list apis to a common interface which makes the SDK simpler.
type ListResponder interface {
	IsListResponder()
}

// GetListResponse returns the ListResponder from the TransactionListResponse.
func (t *TransactionListResponse) GetListResponse() ListResponder {
	switch v := t.Response.(type) {
	case *TransactionListResponse_Result:
		return v.Result
	case *TransactionListResponse_Finished:
		return v.Finished
	}
	return nil
}

// GetListResponse returns the ListResponder from the ListResponse.
func (t *ListResponse) GetListResponse() ListResponder {
	switch v := t.Response.(type) {
	case *ListResponse_Result:
		return v.Result
	case *ListResponse_Finished:
		return v.Finished
	}
	return nil
}
