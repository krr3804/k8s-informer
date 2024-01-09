// Code generated by smithy-go-codegen DO NOT EDIT.

package eks

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the managed node groups associated with the specified cluster in your
// Amazon Web Services account in the specified Amazon Web Services Region.
// Self-managed node groups aren't listed.
func (c *Client) ListNodegroups(ctx context.Context, params *ListNodegroupsInput, optFns ...func(*Options)) (*ListNodegroupsOutput, error) {
	if params == nil {
		params = &ListNodegroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListNodegroups", params, optFns, c.addOperationListNodegroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListNodegroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListNodegroupsInput struct {

	// The name of your cluster.
	//
	// This member is required.
	ClusterName *string

	// The maximum number of results, returned in paginated output. You receive
	// maxResults in a single page, along with a nextToken response element. You can
	// see the remaining results of the initial request by sending another request with
	// the returned nextToken value. This value can be between 1 and 100. If you don't
	// use this parameter, 100 results and a nextToken value, if applicable, are
	// returned.
	MaxResults *int32

	// The nextToken value returned from a previous paginated request, where maxResults
	// was used and the results exceeded the value of that parameter. Pagination
	// continues from the end of the previous results that returned the nextToken
	// value. This value is null when there are no more results to return. This token
	// should be treated as an opaque identifier that is used only to retrieve the next
	// items in a list and not for other programmatic purposes.
	NextToken *string

	noSmithyDocumentSerde
}

type ListNodegroupsOutput struct {

	// The nextToken value returned from a previous paginated request, where maxResults
	// was used and the results exceeded the value of that parameter. Pagination
	// continues from the end of the previous results that returned the nextToken
	// value. This value is null when there are no more results to return. This token
	// should be treated as an opaque identifier that is used only to retrieve the next
	// items in a list and not for other programmatic purposes.
	NextToken *string

	// A list of all of the node groups associated with the specified cluster.
	Nodegroups []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListNodegroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListNodegroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListNodegroups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListNodegroups"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListNodegroupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListNodegroups(options.Region), middleware.Before); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

// ListNodegroupsAPIClient is a client that implements the ListNodegroups
// operation.
type ListNodegroupsAPIClient interface {
	ListNodegroups(context.Context, *ListNodegroupsInput, ...func(*Options)) (*ListNodegroupsOutput, error)
}

var _ ListNodegroupsAPIClient = (*Client)(nil)

// ListNodegroupsPaginatorOptions is the paginator options for ListNodegroups
type ListNodegroupsPaginatorOptions struct {
	// The maximum number of results, returned in paginated output. You receive
	// maxResults in a single page, along with a nextToken response element. You can
	// see the remaining results of the initial request by sending another request with
	// the returned nextToken value. This value can be between 1 and 100. If you don't
	// use this parameter, 100 results and a nextToken value, if applicable, are
	// returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListNodegroupsPaginator is a paginator for ListNodegroups
type ListNodegroupsPaginator struct {
	options   ListNodegroupsPaginatorOptions
	client    ListNodegroupsAPIClient
	params    *ListNodegroupsInput
	nextToken *string
	firstPage bool
}

// NewListNodegroupsPaginator returns a new ListNodegroupsPaginator
func NewListNodegroupsPaginator(client ListNodegroupsAPIClient, params *ListNodegroupsInput, optFns ...func(*ListNodegroupsPaginatorOptions)) *ListNodegroupsPaginator {
	if params == nil {
		params = &ListNodegroupsInput{}
	}

	options := ListNodegroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListNodegroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListNodegroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListNodegroups page.
func (p *ListNodegroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListNodegroupsOutput, error) {
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

	result, err := p.client.ListNodegroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListNodegroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListNodegroups",
	}
}
