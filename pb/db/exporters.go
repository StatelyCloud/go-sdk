package db

// IsItemPayload exports the isItem_Payload type from this package as a convenience for consumers.
type IsItemPayload = isItem_Payload

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
	// Note: make sure to access t.Response via the nil-safe accessor.
	switch v := t.GetResponse().(type) {
	case *TransactionListResponse_Result:
		return v.Result
	case *TransactionListResponse_Finished:
		return v.Finished
	}
	return nil
}

// GetListResponse returns the ListResponder from the ListResponse.
func (t *ListResponse) GetListResponse() ListResponder {
	// Note: make sure to access t.Response via the nil-safe accessor.
	switch v := t.GetResponse().(type) {
	case *ListResponse_Result:
		return v.Result
	case *ListResponse_Finished:
		return v.Finished
	}
	return nil
}

// SyncResponder gets around the unexported interface in the generated proto
// code and allows us to use a common interface for both "Result" and "Finished".
type SyncResponder interface {
	IsSyncResponder()
}

// IsSyncResponder is a type that binds ListPartialResult to the SyncResponder interface.
// This makes handling the response easier in the SDK.
func (*ListFinished) IsSyncResponder() {}

// IsSyncResponder is a type that binds SyncListPartialResponse to the SyncResponder interface.
// This makes handling the response easier in the SDK.
func (*SyncListPartialResponse) IsSyncResponder() {}

// GetSyncResponse marshals the response into a SyncResponder so that we can use
// a common interface for both "Result" and "Finished".
func (t *SyncListResponse) GetSyncResponse() SyncResponder {
	// Note: make sure to access t.Response via the nil-safe accessor.
	switch v := t.GetResponse().(type) {
	case *SyncListResponse_Result:
		return v.Result
	case *SyncListResponse_Finished:
		return v.Finished
	}
	return nil
}
