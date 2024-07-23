package stately

import (
	"github.com/StatelyCloud/go-sdk/pb/db"
)

// ItemTypeMapper is a function that maps a db.Item to your SDK generated types.
// We will generate this type mapper when using Stately's code generation to handle the unmarshalling for you.
type ItemTypeMapper func(item *db.Item) (Item, error)
