package dbmanagement

import "github.com/StatelyCloud/go-sdk/client"

// StoreInfo is information about the store.
type StoreInfo struct {
	ID          client.StoreID
	Name        string
	Description string
}
