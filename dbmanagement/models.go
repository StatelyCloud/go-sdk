package dbmanagement

import (
	"github.com/StatelyCloud/go-sdk/pb/dbmanagement"
	"github.com/StatelyCloud/go-sdk/stately"
)

// StoreInfo is information about the store.
type StoreInfo struct {
	ID            stately.StoreID
	Name          string
	Description   string
	DefaultRegion dbmanagement.Region
	SchemaID      stately.SchemaID
}

// SchemaInfo is information about the schema.
type SchemaInfo struct {
	ID          stately.SchemaID
	Name        string
	Description string
}
