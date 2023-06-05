// Code generated by smithy-go-codegen DO NOT EDIT.

package finspace

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/finspace/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of clusters.
func (c *Client) ListKxClusters(ctx context.Context, params *ListKxClustersInput, optFns ...func(*Options)) (*ListKxClustersOutput, error) {
	if params == nil {
		params = &ListKxClustersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListKxClusters", params, optFns, c.addOperationListKxClustersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListKxClustersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListKxClustersInput struct {

	// A unique identifier for the kdb environment.
	//
	// This member is required.
	EnvironmentId *string

	// Specifies the type of KDB database that is being created. The following types
	// are available:
	//   - HDB – A Historical Database. The data is only accessible with read-only
	//   permissions from one of the FinSpace managed kdb databases mounted to the
	//   cluster.
	//   - RDB – A Realtime Database. This type of database captures all the data from
	//   a ticker plant and stores it in memory until the end of day, after which it
	//   writes all of its data to a disk and reloads the HDB. This cluster type requires
	//   local storage for temporary storage of data during the savedown process. If you
	//   specify this field in your request, you must provide the
	//   savedownStorageConfiguration parameter.
	//   - GATEWAY – A gateway cluster allows you to access data across processes in
	//   kdb systems. It allows you to create your own routing logic using the
	//   initialization scripts and custom code. This type of cluster does not require a
	//   writable local storage.
	ClusterType types.KxClusterType

	// The maximum number of results to return in this request.
	MaxResults int32

	// A token that indicates where a results page should begin.
	NextToken *string

	noSmithyDocumentSerde
}

type ListKxClustersOutput struct {

	// Lists the cluster details.
	KxClusterSummaries []types.KxCluster

	// A token that indicates where a results page should begin.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListKxClustersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListKxClusters{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListKxClusters{}, middleware.After)
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
	if err = addRestJsonContentTypeCustomization(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListKxClustersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListKxClusters(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListKxClusters(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "finspace",
		OperationName: "ListKxClusters",
	}
}
