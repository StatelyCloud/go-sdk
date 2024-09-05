package dbmanagement

import (
	"github.com/StatelyCloud/go-sdk/stately"
)

// StoreInfo is information about the store.
type StoreInfo struct {
	ID          stately.StoreID
	Name        string
	Description string
}
