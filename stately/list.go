package stately

import (
	"context"
	"errors"
	"fmt"

	"connectrpc.com/connect"

	"github.com/StatelyCloud/go-sdk/pb/db"
	"github.com/StatelyCloud/go-sdk/sdkerror"
)

// ListOptions are optional parameters for List.
type ListOptions struct {
	// Limit is the maximum number of items to return. The default is unlimited -
	// all items will be returned.
	Limit uint32

	// SortableProperty is the property to sort by. Default is SortByKeyPath.
	SortableProperty SortableProperty

	// SortDirection is the direction to sort by. Default is Ascending.
	SortDirection SortDirection

	// ItemTypes are a list of item types to include in the result set.
	// If not provided, all item types will be returned.
	ItemTypes []string

	// KeyConditions are additional constraints to apply to the list operation
	// to limit the scope of items to return.
	//
	// At most two KeyConditions can be provided:
	// - one GreaterThan (or GreaterThanOrEqualTo) condition
	// - one LessThan (or LessThanOrEqualTo) condition
	KeyConditions []KeyCondition

	// CelExpressionFilters are CEL expression filters to apply to the result set.
	// Each expression is evaluated on an item type basis, so you can have multiple
	// expressions for different item types, and the existence of a filter for one
	// item type does not mean that other item types are excluded from the result set.
	// To ensure that ONLY specific item types are returned, use the ItemTypes field above.
	CelExpressionFilters []CelExpressionFilter
}

// WithKeyGreaterThan adds a KeyCondition to the ListOptions that restricts the result set
// to items with keys greater than the specified keyPath.
func (lo ListOptions) WithKeyGreaterThan(keyPath string) ListOptions {
	if keyPath == "" {
		return lo // noop on empty keyPath
	}
	lo.KeyConditions = append(lo.KeyConditions, KeyCondition{
		KeyPath:  keyPath,
		Operator: GreaterThan,
	})
	return lo
}

// WithKeyGreaterThanOrEqualTo adds a KeyCondition to the ListOptions that restricts the result set
// to items with keys greater than or equal to the specified keyPath.
func (lo ListOptions) WithKeyGreaterThanOrEqualTo(keyPath string) ListOptions {
	if keyPath == "" {
		return lo // noop on empty keyPath
	}
	lo.KeyConditions = append(lo.KeyConditions, KeyCondition{
		KeyPath:  keyPath,
		Operator: GreaterThanOrEqualTo,
	})
	return lo
}

// WithKeyLessThan adds a KeyCondition to the ListOptions that restricts the result set
// to items with keys less than the specified keyPath.
func (lo ListOptions) WithKeyLessThan(keyPath string) ListOptions {
	if keyPath == "" {
		return lo // noop on empty keyPath
	}
	lo.KeyConditions = append(lo.KeyConditions, KeyCondition{
		KeyPath:  keyPath,
		Operator: LessThan,
	})
	return lo
}

// WithKeyLessThanOrEqualTo adds a KeyCondition to the ListOptions that restricts the result set
// to items with keys less than or equal to the specified keyPath.
func (lo ListOptions) WithKeyLessThanOrEqualTo(keyPath string) ListOptions {
	if keyPath == "" {
		return lo // noop on empty keyPath
	}
	lo.KeyConditions = append(lo.KeyConditions, KeyCondition{
		KeyPath:  keyPath,
		Operator: LessThanOrEqualTo,
	})
	return lo
}

// WithItemTypesToInclude adds ItemType filters to the ListOptions.
func (lo ListOptions) WithItemTypesToInclude(itemTypes ...string) ListOptions {
	lo.ItemTypes = append(lo.ItemTypes, itemTypes...)
	return lo
}

// WithCelExpressionFilter adds a CEL expression filter to the ListOptions.
func (lo ListOptions) WithCelExpressionFilter(itemType, expression string) ListOptions {
	lo.CelExpressionFilters = append(lo.CelExpressionFilters, CelExpressionFilter{
		ItemType:   itemType,
		Expression: expression,
	})
	return lo
}

// WithLimit sets the maximum number of items to return in the ListOptions.
// Note: If no limit is set (or limit is zero) all items will be returned.
func (lo ListOptions) WithLimit(limit uint32) ListOptions {
	lo.Limit = limit
	return lo
}

// WithSortableProperty sets the property to sort by in the ListOptions.
func (lo ListOptions) WithSortableProperty(sortable SortableProperty) ListOptions {
	lo.SortableProperty = sortable
	return lo
}

// WithSortDirection sets the direction to sort by in the ListOptions.
func (lo ListOptions) WithSortDirection(direction SortDirection) ListOptions {
	lo.SortDirection = direction
	return lo
}

// Merge combines two ListOptions into one. "other" takes precedence over "this".
// Nils will overwrite non-nil values.
func (lo *ListOptions) Merge(other *ListOptions) *ListOptions {
	if other == nil {
		return lo
	}
	if lo == nil {
		return other
	}
	lo.Limit = other.Limit
	lo.SortableProperty = other.SortableProperty
	lo.SortDirection = other.SortDirection
	lo.ItemTypes = other.ItemTypes
	lo.KeyConditions = other.KeyConditions
	lo.CelExpressionFilters = other.CelExpressionFilters
	return lo
}

func (lo *ListOptions) filters() []*db.FilterCondition {
	if len(lo.ItemTypes) == 0 && len(lo.CelExpressionFilters) == 0 {
		return nil
	}
	filters := make([]*db.FilterCondition, len(lo.ItemTypes)+len(lo.CelExpressionFilters))
	for i, itemType := range lo.ItemTypes {
		filters[i] = &db.FilterCondition{
			Value: &db.FilterCondition_ItemType{
				ItemType: itemType,
			},
		}
	}
	for i, filter := range lo.CelExpressionFilters {
		filters[len(lo.ItemTypes)+i] = &db.FilterCondition{
			Value: &db.FilterCondition_CelExpression{
				CelExpression: filter.toProto(),
			},
		}
	}
	return filters
}

func (lo *ListOptions) keyConditions() ([]*db.KeyCondition, error) {
	if len(lo.KeyConditions) == 0 {
		return nil, nil
	}

	result := make([]*db.KeyCondition, len(lo.KeyConditions))
	var err error
	for i, cond := range lo.KeyConditions {
		result[i], err = cond.toProto()
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// ContinueOptions are optional parameters for Continue.
type ContinueOptions struct {
	// SortDirection is the direction to sort by. Default is Ascending.
	SortDirection SortDirection
}

// SortableProperty is the property to sort by.
type SortableProperty int32

const (
	// SortByKeyPath sorts by the key path.
	SortByKeyPath SortableProperty = iota
	// SortByLastModifiedVersion sorts by the last time the item was modified.
	SortByLastModifiedVersion
)

// SortDirection is the direction to sort by.
type SortDirection int32

const (
	// Ascending is the default sort direction.
	Ascending SortDirection = iota
	// Descending is the reverse sort direction.
	Descending
)

// KeyConditionOperator defines the operator to use for key conditions.
type KeyConditionOperator int32

const (
	// GreaterThan specifies that a key must be greater than the specified value to be
	// included in the result set.
	// In ascending order, this can also be thought of as a "start" condition.
	// In descending order, this can be thought of as an "end" condition.
	GreaterThan KeyConditionOperator = iota + 1

	// GreaterThanOrEqualTo specifies that a key must be greater than or equal to the
	// specified value to be included in the result set.
	// In ascending order, this can also be thought of as a "start" condition.
	// In descending order, this can be thought of as an "end" condition.
	GreaterThanOrEqualTo

	// LessThan specifies that a key must be less than the specified value to be
	// included in the result set.
	// In ascending order, this can also be thought of as an "end" condition.
	// In descending order, this can be thought of as a "start" condition.
	LessThan

	// LessThanOrEqualTo specifies that a key must be less than or equal to the
	// specified value to be included in the result set.
	// In ascending order, this can also be thought of as an "end" condition.
	// In descending order, this can be thought of as a "start" condition.
	LessThanOrEqualTo
)

// A KeyCondition is an additional constraint to apply to a list operation to limit
// the scope of items to return.
// Stately applies these conditions at the DB layer to optimize the list operation
// latency and cost.
type KeyCondition struct {
	// KeyPath is a valid key prefix (or full key) used to filter or optimize the
	// list operation based on the operator below.
	//
	// Note: When using KeyConditions and KeyPrefixes together, KeyCondition KeyPaths must
	// share the same prefix. For example a KeyCondition:
	//
	//	{ Operator: GreaterThan KeyPath: "/group-MY_GROUP_ID/category-10/item-10" }
	//
	// Can be used with any of the key prefixes:
	// - /group-MY_GROUP_ID
	// - /group-MY_GROUP_ID/category
	// - /group-MY_GROUP_ID/category-10
	// - /group-MY_GROUP_ID/category-10/item
	//
	// But cannot be used with the key prefixes:
	// - /group-MY_GROUP_ID/category-11
	// - /group-MY_GROUP_ID/otherSubgroup
	KeyPath string

	// Operator specifies how to apply a KeyPath condition to the list operation.
	// Valid operators are:
	// - GreaterThan: Items returned must have a key greater than the key path above.
	// - GreaterThanOrEqualTo: Items returned must have a key greater than or equal to the key path above.
	// - LessThan: Items returned must have a key less than the key path above.
	// - LessThanOrEqualTo: Items returned must have a key less than or equal to the key path above.
	Operator KeyConditionOperator
}

func (kc KeyCondition) toProto() (*db.KeyCondition, error) {
	result := &db.KeyCondition{
		KeyPath: kc.KeyPath,
	}

	switch kc.Operator {
	case GreaterThan:
		result.Operator = db.Operator_OPERATOR_GREATER_THAN
	case GreaterThanOrEqualTo:
		result.Operator = db.Operator_OPERATOR_GREATER_THAN_OR_EQUAL
	case LessThan:
		result.Operator = db.Operator_OPERATOR_LESS_THAN
	case LessThanOrEqualTo:
		result.Operator = db.Operator_OPERATOR_LESS_THAN_OR_EQUAL
	default:
		return nil, &sdkerror.Error{
			Code:        connect.CodeInvalidArgument,
			StatelyCode: "InvalidKeyConditionOperator",
			Message:     fmt.Sprintf("invalid key condition operator: %d", kc.Operator),
			CauseErr:    nil,
		}
	}

	return result, nil
}

type CelExpressionFilter struct {
	// ItemType is the ItemType the filter applies to.
	// Note: This filter has no effect on other ItemTypes that may be listed over; they
	// will still be included in the results unless other filters are applied to them.
	ItemType string

	// Expression is the CEL expression to evaluate for the ItemType above.
	// If the expression evaluates to true, the item is included in the results,
	// otherwise it is excluded.
	//
	// In the context of the CEL expression, 'this' refers to the item being evaluated,
	// and properties should be accessed by the names are as they appear in schema
	// (this is also the json field name which can be found in the json tags of
	// the generated schema) not in generated code name. For example, the following
	// item type defined in schema:
	//
	// 	itemType("Person", {
	// 	 keyPath: "/Person-:id",
	// 	 fields: {
	// 	   id: { type: uuid, initialValue: "uuid" },
	// 	   full_name: { type: string },
	// 	   age: { type: int, required: false },
	// 	   email: { type: string, required: false },
	// 	 }
	// 	})
	//
	// in generated code looks like:
	//
	// 	type Person struct {
	//		Id uuid.UUID `protobuf:"bytes,1" json:"id,omitempty"`
	//		FullName string `protobuf:"bytes,2" json:"full_name,omitempty"`
	//		Age int64 `protobuf:"zigzag64,3" json:"age,omitempty,string"`
	//		Email string `protobuf:"bytes,4" json:"email,omitempty"`
	// 	}
	//
	// So, we could build a CEL expression like:
	//
	// 	this.full_name.contains('John') && this.age > 30 && !has(this.email)
	//
	// This will include all Person items where the full_name contains 'John'
	// and the age is greater than 30, and the email field is absent.
	// For more about CEL expressions, see the CEL documentation:
	// https://github.com/google/cel-spec/blob/master/doc/langdef.md
	Expression string
}

func (cef CelExpressionFilter) toProto() *db.CelExpression {
	return &db.CelExpression{
		ItemType:   cef.ItemType,
		Expression: cef.Expression,
	}
}

// ListToken is a stateless token that saves your place in a result set,
// allowing you to fetch additional results with ContinueList, or get updated
// results with SyncList.
type ListToken struct {
	// Data will never be nil. This is the token data that you pass to
	// ContinueList or SyncList.
	Data []byte
	// CanContinue indicates if there are more results to fetch using
	// ContinueList.
	CanContinue bool
	// CanSync indicates that you could call SyncList with this token later to get
	// updated items. This is determined by the type of store you're listing from.
	CanSync bool

	// SchemaVersionID is the schema version ID of the store that produced this token.
	// When making ContinueList calls, ensure your client version uses tokens that
	// match this field. Using tokens with different schema versions will result
	// in a SchemaVersionMismatch error. For SyncList calls, you only need to
	// ensure you handle Reset responses correctly.
	SchemaVersionID SchemaVersionID
}

// newToken creates a new ListToken from a proto token and a store.
// If the proto token is nil, newToken will return nil.
func newToken(token *db.ListToken) *ListToken {
	if token == nil {
		return nil
	}
	return &ListToken{
		Data:            token.GetTokenData(),
		CanContinue:     token.GetCanContinue(),
		CanSync:         token.GetCanSync(),
		SchemaVersionID: SchemaVersionID(token.GetSchemaVersionId()),
	}
}

// stream enables us to abstract txn and non-txn streams. For txn we need to
// verify messageId but for non-txn we don't.
type stream struct {
	// receive is a function that reads the next item from the stream. It
	// should return false if the stream is done.
	receive  func() bool
	response db.ListResponder
	err      error
}

type listIterator struct {
	stream     *stream
	itemMapper ItemTypeMapper

	// tracks where we're at in the current response page
	currPos  int
	currResp []*db.Item
	currItem Item

	// holds the final token and error to be returned by Token()
	finalToken *ListToken
	finalErr   error
}

// Next reads an item of the stream, and populates Item() with the current item.
// If there are no more items OR there is an error, Next will return false and
// the error will be available via Token().
func (li *listIterator) Next() bool {
	// if we're beyond the current response page, get the next page
	if li.currPos >= len(li.currResp) {

		// if the stream is done, return any possible error from the stream
		if !li.stream.receive() {
			li.finalErr = li.stream.err
			return false
		}

		switch v := li.stream.response.(type) {
		case *db.ListPartialResult:
			li.currResp = v.GetItems()
			li.currPos = 0
		case *db.ListFinished:
			li.finalToken = newToken(v.GetToken())
		}

	}
	// If we see the final token, we should stop
	if li.finalToken != nil {
		return false
	}

	// if we have items in the current response page, return the next one
	li.currItem, li.finalErr = li.itemMapper(li.currResp[li.currPos])
	li.currPos++

	// if there's no error, continue iterating
	return li.finalErr == nil
}

// Token returns the current token OR any error that occurred during iteration.
func (li *listIterator) Token() (*ListToken, error) {
	return li.finalToken, li.finalErr
}

// Value returns the current item in the iteration.
func (li *listIterator) Value() Item {
	return li.currItem
}

func (c *client) ContinueList(ctx context.Context, token []byte) (ListResponse[Item], error) {
	if token == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("token is nil"))
	}

	// call continue list
	response, err := c.client.ContinueList(ctx, connect.NewRequest(&db.ContinueListRequest{
		TokenData:       token,
		SchemaVersionId: uint32(c.schemaVersionID),
	}))
	if err != nil {
		return nil, err
	}

	return &listIterator{
		stream:     newStream(response),
		itemMapper: c.itemMapper,
	}, nil
}

func (c *client) BeginList(
	ctx context.Context,
	keyPath string,
	opts ...ListOptions,
) (ListResponse[Item], error) {
	options := &ListOptions{}
	for _, opt := range opts {
		options = options.Merge(&opt)
	}

	keyConditions, err := options.keyConditions()
	if err != nil {
		return nil, err
	}

	response, err := c.client.BeginList(ctx, connect.NewRequest(&db.BeginListRequest{
		StoreId:          uint64(c.storeID),
		SchemaVersionId:  uint32(c.schemaVersionID),
		KeyPathPrefix:    keyPath,
		AllowStale:       c.allowStale,
		Limit:            options.Limit,
		SortProperty:     db.SortableProperty(options.SortableProperty),
		SortDirection:    db.SortDirection(options.SortDirection),
		FilterConditions: options.filters(),
		KeyConditions:    keyConditions,
	}))
	if err != nil {
		return nil, err
	}

	return &listIterator{
		stream:     newStream(response),
		itemMapper: c.itemMapper,
	}, nil
}

func newStream(response *connect.ServerStreamForClient[db.ListResponse]) *stream {
	newStream := &stream{}

	newStream.receive = func() bool {
		canContinue := response.Receive()
		newStream.response = response.Msg().GetListResponse()
		newStream.err = response.Err()
		return canContinue
	}
	return newStream
}
