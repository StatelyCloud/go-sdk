package data

import (
	"errors"

	"connectrpc.com/connect"

	pb "github.com/StatelyCloud/go-sdk/pb/data"
)

// Get retrieves an item from the store.
func (t *transaction) Get(item string) (*RawItem, error) {
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
func (t *transaction) GetBatch(items ...string) ([]*RawItem, error) {
	req := t.newTXNReq(&pb.TransactionRequest_GetItems{
		GetItems: &pb.TransactionGet{Gets: mapToItemKey(items)},
	})
	err := t.stream.Send(req)
	if err != nil {
		return nil, err
	}

	res, err := receiveExpected(t, req.GetMessageId(), (*pb.TransactionResponse).GetGetResults)
	if err != nil {
		return nil, err
	}

	results := make([]*RawItem, len(res.Items))
	for idx, v := range res.Items {
		item, err := protoToItem(v)
		if err != nil {
			return nil, err
		}
		results[idx] = item
	}

	return results, nil
}

// Put is a convenience method for adding a single Item to the Store, or replacing the RawItem if it exists at that path.
func (t *transaction) Put(path string, data any) error {
	return t.PutBatch(&PutData{KeyPath: path, Data: data})
}

// PutBatch schedules items to be written with new keys on commit.
func (t *transaction) PutBatch(items ...*PutData) error {
	putItems, originalItems, err := mapPutRequest(items)
	if err != nil {
		return err
	}

	err = t.stream.Send(t.newTXNReq(&pb.TransactionRequest_PutItems{
		PutItems: &pb.TransactionPut{Puts: putItems},
	}))
	if err != nil {
		return err
	}

	t.putRequests = append(t.putRequests, originalItems...)

	return nil
}

// Append adds one Item to a parent path, automatically assigning an ID.
func (t *transaction) Append(
	prefix string,
	itemType string,
	data any,
	idAssignment AppendIDAssignment,
) (string, error) {
	paths, err := t.AppendBatch(prefix, &AppendRequest{ItemType: itemType, Data: data, IDAssignment: idAssignment})
	if err != nil {
		return "", err
	}
	if len(paths) == 0 {
		return "", nil
	}
	return paths[0], nil
}

// AppendBatch schedules items to be appended with new keys on commit.
func (t *transaction) AppendBatch(prefix string, items ...*AppendRequest) ([]string, error) {
	requests, requestData, err := mapAppendRequest(items)
	if err != nil {
		return nil, err
	}

	req := t.newTXNReq(&pb.TransactionRequest_AppendItems{
		AppendItems: &pb.TransactionAppend{
			ParentPath: prefix,
			Appends:    requests,
		},
	})
	err = t.stream.Send(req)
	if err != nil {
		return nil, err
	}

	res, err := receiveExpected(t, req.GetMessageId(), (*pb.TransactionResponse).GetAppendAck)
	if err != nil {
		return nil, err
	}

	t.appendRequests = append(t.appendRequests, requestData...)

	return res.KeyPaths, nil
}

// Delete schedules items to be deleted on commit.
func (t *transaction) Delete(items ...string) error {
	err := t.stream.Send(t.newTXNReq(&pb.TransactionRequest_DeleteItems{
		DeleteItems: &pb.TransactionDelete{Deletes: mapDeleteRequest(items)},
	}))
	if err != nil {
		return err
	}

	return nil
}

// BeginList is like a query only we call it 'List'.
func (t *transaction) BeginList(prefix string, options ...*ListOptions) (ListResponse[*RawItem], error) {
	opts := ListOptions{}
	for _, opt := range options {
		opts.Merge(opt)
	}

	req := t.newTXNReq(&pb.TransactionRequest_BeginList{
		BeginList: &pb.TransactionBeginList{
			KeyPathPrefix: prefix,
			Limit:         opts.Limit,
			SortProperty:  pb.SortableProperty(opts.SortableProperty),
			SortDirection: pb.SortDirection(opts.SortDirection),
		},
	})
	err := t.stream.Send(req)
	if err != nil {
		return nil, err
	}

	return &listIterator{
		stream: t.newListStream(req.GetMessageId()),
	}, nil
}

func (t *transaction) ContinueList(token *ListToken) (ListResponse[*RawItem], error) {
	if token == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("token is nil"))
	}
	req := t.newTXNReq(&pb.TransactionRequest_ContinueList{
		ContinueList: &pb.TransactionContinueList{
			TokenData: token.Data,
		},
	})
	err := t.stream.Send(req)
	if err != nil {
		return nil, err
	}

	return &listIterator{
		stream: t.newListStream(req.GetMessageId()),
	}, nil
}

// newTXNReq converts a transaction command to a transaction request.
func (t *transaction) newTXNReq(command pb.IsTransactionCommand) *pb.TransactionRequest {
	return &pb.TransactionRequest{
		MessageId: t.id.Add(1), // increment the message ID
		Command:   command,
	}
}

func (t *transaction) newListStream(msgID uint32) *stream {
	newStream := &stream{}

	// pull a message off the txn stream, parse and set it to resp or err
	newStream.receive = func() bool {
		var res *pb.TransactionListResponse
		res, newStream.err = receiveExpected(t, msgID, (*pb.TransactionResponse).GetListResults)
		newStream.response = res.GetListResponse()
		return !t.done.Load()
	}

	return newStream
}
