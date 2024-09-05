package stately

import (
	"errors"

	"connectrpc.com/connect"

	"github.com/StatelyCloud/go-sdk/pb/db"
)

// Get retrieves an item from the store.
func (t *transaction) Get(item string) (Item, error) {
	items, err := t.GetBatch(item)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, nil
	}
	return items[0], nil
}

// GetBatch retrieves items from the store.
func (t *transaction) GetBatch(itemKeys ...string) ([]Item, error) {
	req := t.newTXNReq(&db.TransactionRequest_GetItems{
		GetItems: &db.TransactionGet{Gets: mapToItemKey(itemKeys)},
	})
	err := t.safeSend(req)
	if err != nil {
		return nil, err
	}

	res, err := receiveExpected(t, req.GetMessageId(), (*db.TransactionResponse).GetGetResults)
	if err != nil {
		return nil, err
	}

	results := make([]Item, len(res.Items))
	for idx, v := range res.Items {
		item, err := t.itemMapper(v)
		if err != nil {
			return nil, err
		}
		results[idx] = item
	}

	return results, nil
}

// Put is a convenience method for adding a single Item to the Store, or
// replacing the RawItem if it exists at that path.
func (t *transaction) Put(item Item) (GeneratedID, error) {
	items, err := t.PutBatch(item)
	if err != nil || len(items) == 0 {
		return GeneratedID{}, err
	}
	return items[0], nil
}

// PutBatch schedules items to be written with new keys on commit.
func (t *transaction) PutBatch(items ...Item) ([]GeneratedID, error) {
	putItems, err := mapPutRequest(items)
	if err != nil {
		return nil, err
	}

	req := t.newTXNReq(&db.TransactionRequest_PutItems{
		PutItems: &db.TransactionPut{Puts: putItems},
	})

	err = t.safeSend(req)
	if err != nil {
		return nil, err
	}

	res, err := receiveExpected(t, req.GetMessageId(), (*db.TransactionResponse).GetPutAck)
	if err != nil {
		return nil, err
	}
	t.putRequests = append(t.putRequests, items...)

	// map the results back
	generatedIDs := make([]GeneratedID, len(res.GeneratedIds))
	for idx, v := range res.GeneratedIds {
		var generatedID GeneratedID
		switch v.GetValue().(type) {
		case *db.GeneratedID_Uint:
			generatedID = GeneratedID{
				Uint64: v.GetUint(),
			}
		case *db.GeneratedID_Bytes:
			generatedID = GeneratedID{
				Bytes: v.GetBytes(),
			}
		}

		generatedIDs[idx] = generatedID
	}

	return generatedIDs, nil
}

// Delete schedules items to be deleted on commit.
func (t *transaction) Delete(itemKeys ...string) error {
	err := t.safeSend(t.newTXNReq(&db.TransactionRequest_DeleteItems{
		DeleteItems: &db.TransactionDelete{Deletes: mapDeleteRequest(itemKeys)},
	}))
	if err != nil {
		return err
	}

	return nil
}

// BeginList is like a query only we call it 'List'.
func (t *transaction) BeginList(prefix string, options ...ListOptions) (ListResponse[Item], error) {
	opts := ListOptions{}
	for _, opt := range options {
		opts.Merge(&opt)
	}

	req := t.newTXNReq(&db.TransactionRequest_BeginList{
		BeginList: &db.TransactionBeginList{
			KeyPathPrefix: prefix,
			Limit:         opts.Limit,
			SortProperty:  db.SortableProperty(opts.SortableProperty),
			SortDirection: db.SortDirection(opts.SortDirection),
		},
	})
	err := t.safeSend(req)
	if err != nil {
		return nil, err
	}

	return &listIterator{
		stream:     t.newListStream(req.GetMessageId()),
		itemMapper: t.itemMapper,
	}, nil
}

func (t *transaction) ContinueList(token *ListToken) (ListResponse[Item], error) {
	if token == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("token is nil"))
	}
	req := t.newTXNReq(&db.TransactionRequest_ContinueList{
		ContinueList: &db.TransactionContinueList{
			TokenData: token.Data,
		},
	})
	err := t.safeSend(req)
	if err != nil {
		return nil, err
	}

	return &listIterator{
		stream:     t.newListStream(req.GetMessageId()),
		itemMapper: t.itemMapper,
	}, nil
}

// newTXNReq converts a transaction command to a transaction request.
func (t *transaction) newTXNReq(command db.IsTransactionCommand) *db.TransactionRequest {
	return &db.TransactionRequest{
		MessageId: t.id.Add(1), // increment the message ID
		Command:   command,
	}
}

func (t *transaction) newListStream(msgID uint32) *stream {
	newStream := &stream{}

	// pull a message off the txn stream, parse and set it to resp or err
	newStream.receive = func() bool {
		var res *db.TransactionListResponse
		res, newStream.err = receiveExpected(t, msgID, (*db.TransactionResponse).GetListResults)
		newStream.response = res.GetListResponse()
		return !t.done.Load()
	}

	return newStream
}
