// Code generated by smithy-go-codegen DO NOT EDIT.

package internetmonitor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/internetmonitor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all of your monitors for Amazon CloudWatch Internet Monitor and their
// statuses, along with the Amazon Resource Name (ARN) and name of each monitor.
func (c *Client) ListMonitors(ctx context.Context, params *ListMonitorsInput, optFns ...func(*Options)) (*ListMonitorsOutput, error) {
	if params == nil {
		params = &ListMonitorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMonitors", params, optFns, c.addOperationListMonitorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMonitorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMonitorsInput struct {

	// A boolean option that you can set to TRUE to include monitors for linked
	// accounts in a list of monitors, when you've set up cross-account sharing in
	// Amazon CloudWatch Internet Monitor. You configure cross-account sharing by using
	// Amazon CloudWatch Observability Access Manager. For more information, see
	// Internet Monitor cross-account observability (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cwim-cross-account.html)
	// in the Amazon CloudWatch Internet Monitor User Guide.
	IncludeLinkedAccounts *bool

	// The number of monitor objects that you want to return with this call.
	MaxResults *int32

	// The status of a monitor. This includes the status of the data processing for
	// the monitor and the status of the monitor itself. For information about the
	// statuses for a monitor, see Monitor (https://docs.aws.amazon.com/internet-monitor/latest/api/API_Monitor.html)
	// .
	MonitorStatus *string

	// The token for the next set of results. You receive this token from a previous
	// call.
	NextToken *string

	noSmithyDocumentSerde
}

type ListMonitorsOutput struct {

	// A list of monitors.
	//
	// This member is required.
	Monitors []types.Monitor

	// The token for the next set of results. You receive this token from a previous
	// call.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListMonitorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListMonitors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListMonitors{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListMonitors"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMonitors(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
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

// ListMonitorsAPIClient is a client that implements the ListMonitors operation.
type ListMonitorsAPIClient interface {
	ListMonitors(context.Context, *ListMonitorsInput, ...func(*Options)) (*ListMonitorsOutput, error)
}

var _ ListMonitorsAPIClient = (*Client)(nil)

// ListMonitorsPaginatorOptions is the paginator options for ListMonitors
type ListMonitorsPaginatorOptions struct {
	// The number of monitor objects that you want to return with this call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListMonitorsPaginator is a paginator for ListMonitors
type ListMonitorsPaginator struct {
	options   ListMonitorsPaginatorOptions
	client    ListMonitorsAPIClient
	params    *ListMonitorsInput
	nextToken *string
	firstPage bool
}

// NewListMonitorsPaginator returns a new ListMonitorsPaginator
func NewListMonitorsPaginator(client ListMonitorsAPIClient, params *ListMonitorsInput, optFns ...func(*ListMonitorsPaginatorOptions)) *ListMonitorsPaginator {
	if params == nil {
		params = &ListMonitorsInput{}
	}

	options := ListMonitorsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListMonitorsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListMonitorsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListMonitors page.
func (p *ListMonitorsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListMonitorsOutput, error) {
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

	result, err := p.client.ListMonitors(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListMonitors(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListMonitors",
	}
}
