// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the historical values for one or more asset properties. For more
// information, see Querying historical values (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/query-industrial-data.html#historical-values)
// in the IoT SiteWise User Guide.
func (c *Client) BatchGetAssetPropertyValueHistory(ctx context.Context, params *BatchGetAssetPropertyValueHistoryInput, optFns ...func(*Options)) (*BatchGetAssetPropertyValueHistoryOutput, error) {
	if params == nil {
		params = &BatchGetAssetPropertyValueHistoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchGetAssetPropertyValueHistory", params, optFns, c.addOperationBatchGetAssetPropertyValueHistoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchGetAssetPropertyValueHistoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchGetAssetPropertyValueHistoryInput struct {

	// The list of asset property historical value entries for the batch get request.
	// You can specify up to 16 entries per request.
	//
	// This member is required.
	Entries []types.BatchGetAssetPropertyValueHistoryEntry

	// The maximum number of results to return for each paginated request. A result
	// set is returned in the two cases, whichever occurs first.
	//   - The size of the result set is equal to 4 MB.
	//   - The number of data points in the result set is equal to the value of
	//   maxResults . The maximum value of maxResults is 20000.
	MaxResults *int32

	// The token to be used for the next set of paginated results.
	NextToken *string

	noSmithyDocumentSerde
}

type BatchGetAssetPropertyValueHistoryOutput struct {

	// A list of the errors (if any) associated with the batch request. Each error
	// entry contains the entryId of the entry that failed.
	//
	// This member is required.
	ErrorEntries []types.BatchGetAssetPropertyValueHistoryErrorEntry

	// A list of entries that were not processed by this batch request. because these
	// entries had been completely processed by previous paginated requests. Each
	// skipped entry contains the entryId of the entry that skipped.
	//
	// This member is required.
	SkippedEntries []types.BatchGetAssetPropertyValueHistorySkippedEntry

	// A list of entries that were processed successfully by this batch request. Each
	// success entry contains the entryId of the entry that succeeded and the latest
	// query result.
	//
	// This member is required.
	SuccessEntries []types.BatchGetAssetPropertyValueHistorySuccessEntry

	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchGetAssetPropertyValueHistoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchGetAssetPropertyValueHistory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchGetAssetPropertyValueHistory{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addEndpointPrefix_opBatchGetAssetPropertyValueHistoryMiddleware(stack); err != nil {
		return err
	}
	if err = addOpBatchGetAssetPropertyValueHistoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetAssetPropertyValueHistory(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

type endpointPrefix_opBatchGetAssetPropertyValueHistoryMiddleware struct {
}

func (*endpointPrefix_opBatchGetAssetPropertyValueHistoryMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opBatchGetAssetPropertyValueHistoryMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "data." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opBatchGetAssetPropertyValueHistoryMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opBatchGetAssetPropertyValueHistoryMiddleware{}, `OperationSerializer`, middleware.After)
}

// BatchGetAssetPropertyValueHistoryAPIClient is a client that implements the
// BatchGetAssetPropertyValueHistory operation.
type BatchGetAssetPropertyValueHistoryAPIClient interface {
	BatchGetAssetPropertyValueHistory(context.Context, *BatchGetAssetPropertyValueHistoryInput, ...func(*Options)) (*BatchGetAssetPropertyValueHistoryOutput, error)
}

var _ BatchGetAssetPropertyValueHistoryAPIClient = (*Client)(nil)

// BatchGetAssetPropertyValueHistoryPaginatorOptions is the paginator options for
// BatchGetAssetPropertyValueHistory
type BatchGetAssetPropertyValueHistoryPaginatorOptions struct {
	// The maximum number of results to return for each paginated request. A result
	// set is returned in the two cases, whichever occurs first.
	//   - The size of the result set is equal to 4 MB.
	//   - The number of data points in the result set is equal to the value of
	//   maxResults . The maximum value of maxResults is 20000.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// BatchGetAssetPropertyValueHistoryPaginator is a paginator for
// BatchGetAssetPropertyValueHistory
type BatchGetAssetPropertyValueHistoryPaginator struct {
	options   BatchGetAssetPropertyValueHistoryPaginatorOptions
	client    BatchGetAssetPropertyValueHistoryAPIClient
	params    *BatchGetAssetPropertyValueHistoryInput
	nextToken *string
	firstPage bool
}

// NewBatchGetAssetPropertyValueHistoryPaginator returns a new
// BatchGetAssetPropertyValueHistoryPaginator
func NewBatchGetAssetPropertyValueHistoryPaginator(client BatchGetAssetPropertyValueHistoryAPIClient, params *BatchGetAssetPropertyValueHistoryInput, optFns ...func(*BatchGetAssetPropertyValueHistoryPaginatorOptions)) *BatchGetAssetPropertyValueHistoryPaginator {
	if params == nil {
		params = &BatchGetAssetPropertyValueHistoryInput{}
	}

	options := BatchGetAssetPropertyValueHistoryPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &BatchGetAssetPropertyValueHistoryPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *BatchGetAssetPropertyValueHistoryPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next BatchGetAssetPropertyValueHistory page.
func (p *BatchGetAssetPropertyValueHistoryPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*BatchGetAssetPropertyValueHistoryOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.BatchGetAssetPropertyValueHistory(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opBatchGetAssetPropertyValueHistory(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "BatchGetAssetPropertyValueHistory",
	}
}
