package stately

import (
	"context"
	"errors"

	"connectrpc.com/connect"

	"github.com/StatelyCloud/go-sdk/pb/db"
)

// ScanOptions are optional parameters for Scan.
type ScanOptions struct {
	// Limit is the maximum number of items to return. If set to 0 then the first
	// page of results will be returned which may empty because it does not
	// contain items of your selected item types. Be sure to check
	// token.canContinue to see if there are more results to fetch. The default 0.
	Limit uint32

	// ItemTypes are the item types to filter by. If not provided, all item
	// types will be returned.
	ItemTypes []string

	// CelExpressionFilters are CEL expression filters to apply to the result set.
	// Each expression is evaluated on an item type basis, so you can have multiple
	// expressions for different item types, and the existence of a filter for one
	// item type does not mean that other item types are excluded from the result set.
	// To ensure that ONLY specific item types are returned, use the ItemTypes field above.
	CelExpressionFilters []CelExpressionFilter

	// TotalSegments is the total number of segments to divide the scan into.
	// If this is provided, then segmentation will be enabled for the scan
	// and this scan will only return items from the segment specified by
	// SegmentIndex. If this is not provided, then segmentation will be
	// disabled and SegmentIndex will be ignored.
	TotalSegments uint32

	// SegmentIndex is the index of the segment to scan. If TotalSegments is
	// provided, SegmentIndex must also be provided and must be less than
	// TotalSegments. If TotalSegments is not provided, SegmentIndex is ignored.
	SegmentIndex uint32
}

// Merge combines two ScanOptions into one. "other" takes precedence over "this".
// Nils will overwrite non-nil values.
func (lo *ScanOptions) Merge(other *ScanOptions) *ScanOptions {
	if other == nil {
		return lo
	}
	if lo == nil {
		return other
	}
	lo.Limit = other.Limit
	lo.ItemTypes = other.ItemTypes
	lo.TotalSegments = other.TotalSegments
	lo.SegmentIndex = other.SegmentIndex
	lo.CelExpressionFilters = other.CelExpressionFilters
	return lo
}

// WithItemTypesToInclude adds ItemType filters to the ScanOptions.
func (lo ScanOptions) WithItemTypesToInclude(itemTypes ...string) ScanOptions {
	lo.ItemTypes = append(lo.ItemTypes, itemTypes...)
	return lo
}

// WithCelExpressionFilter adds a CEL expression filter to the ScanOptions.
func (lo ScanOptions) WithCelExpressionFilter(itemType, expression string) ScanOptions {
	lo.CelExpressionFilters = append(lo.CelExpressionFilters, CelExpressionFilter{
		ItemType:   itemType,
		Expression: expression,
	})
	return lo
}

func (c *client) BeginScan(
	ctx context.Context,
	opts ...ScanOptions,
) (ListResponse[Item], error) {
	options := &ScanOptions{}
	for _, opt := range opts {
		options = options.Merge(&opt)
	}

	var segmentationParams *db.SegmentationParams
	if options.TotalSegments > 0 {
		segmentationParams = &db.SegmentationParams{
			TotalSegments: options.TotalSegments,
			SegmentIndex:  options.SegmentIndex,
		}
	}

	response, err := c.client.BeginScan(ctx, connect.NewRequest(&db.BeginScanRequest{
		StoreId:            uint64(c.storeID),
		SegmentationParams: segmentationParams,
		SchemaVersionId:    uint32(c.schemaVersionID),
		FilterConditions:   buildFilters(options.ItemTypes, options.CelExpressionFilters),
		Limit:              options.Limit,
	}))
	if err != nil {
		return nil, err
	}

	return &listIterator{
		stream:     newStream(response),
		itemMapper: c.itemMapper,
	}, nil
}

func (c *client) ContinueScan(ctx context.Context, token []byte) (ListResponse[Item], error) {
	if token == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("token is nil"))
	}

	// call continue scan
	response, err := c.client.ContinueScan(ctx, connect.NewRequest(&db.ContinueScanRequest{
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
