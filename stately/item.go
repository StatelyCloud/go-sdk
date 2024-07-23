package stately

import (
	"github.com/StatelyCloud/go-sdk/pb/db"
)

// UnknownItemTypeError is returned when we receive an item type in the wire format and don't have a corresponding
// SDK type to map to. This can happen if you're using an older version of the SDK that doesn't understand that type.
type UnknownItemTypeError struct {
	ItemType string
}

func (e UnknownItemTypeError) Error() string {
	return "Unknown item type: " + e.ItemType
}

// Item is the interface that all Stately items implement. We use this interface to distinguish between items that
// can be marshalled and unmarshalled to and from the Stately wire format. Item also exposes functionality for custom
// Marshalling and Unmarshalling of items.
type Item interface {
	// StatelyItemType can be used when switching or mapping item types to determine the type of the item.
	// This is used by the SDK to determine the type of the item. You can also use this if you want to interact with
	// the raw Stately APIs by providing this item type to the relevant Stately APIs
	StatelyItemType() string

	// UnmarshalStately unmarshals the wire format of your db.Item into your SDK generated Item. Invoking this method
	// will overwrite the current state of the underlying Item with he contents of db.Item.
	// If you're using the SDK, you shouldn't need to use this. If you wish to interact with the raw Stately APIs, you
	// can use this method to unmarshal items like so:
	//
	//	item := &myschema.MyItem{}
	//	dbItem, err := item.MarshalStately()
	//	// check for errors
	//	req := &db.PutRequest{
	//		Puts: []*db.PutItem{
	//			{
	//				Item: dbItem,
	//			},
	//		},
	//	}
	//	var client dbconnect.DatabaseServiceClient
	//	// init client
	//	resp, err := client.Put(ctx, connect.NewRequest[db.PutRequest](req))
	//	// check for errors
	//	for _, dbItem := range resp.Msg.Items {
	//		// handle response
	//	}
	UnmarshalStately(item *db.Item) error

	// MarshalStately marshals the contents of your generated SDK Item into our wire format.
	// If you're using the SDK, you shouldn't need to use this. If you wish to interact with the raw Stately APIs, you
	// can use this method to marshal items like so:
	//
	//	req := &db.GetRequest{
	//		Gets: []*db.GetItem{
	//			{KeyPath: "/key-id/"},
	//		},
	//	}
	//	var client dbconnect.DatabaseServiceClient
	//	// init client
	//	resp, err := client.Get(ctx, connect.NewRequest[db.GetRequest](req))
	//	// check for errors
	//	for _, dbItem := range resp.Msg.Items {
	//		item := myschema.MyItem{}
	//		err = item.UnmarshalStately(dbItem)
	//		// check for errors + handle item
	//	}
	MarshalStately() (*db.Item, error)
}
