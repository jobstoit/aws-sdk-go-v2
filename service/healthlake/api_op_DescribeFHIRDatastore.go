// Code generated by smithy-go-codegen DO NOT EDIT.

package healthlake

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/healthlake/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the properties associated with the FHIR data store, including the data
// store ID, data store ARN, data store name, data store status, when the data
// store was created, data store type version, and the data store's endpoint.
func (c *Client) DescribeFHIRDatastore(ctx context.Context, params *DescribeFHIRDatastoreInput, optFns ...func(*Options)) (*DescribeFHIRDatastoreOutput, error) {
	if params == nil {
		params = &DescribeFHIRDatastoreInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeFHIRDatastore", params, optFns, c.addOperationDescribeFHIRDatastoreMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeFHIRDatastoreOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeFHIRDatastoreInput struct {

	// The AWS-generated data store ID.
	//
	// This member is required.
	DatastoreId *string

	noSmithyDocumentSerde
}

type DescribeFHIRDatastoreOutput struct {

	// All properties associated with a data store, including the data store ID, data
	// store ARN, data store name, data store status, when the data store was created,
	// data store type version, and the data store's endpoint.
	//
	// This member is required.
	DatastoreProperties *types.DatastoreProperties

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeFHIRDatastoreMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeFHIRDatastore{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeFHIRDatastore{}, middleware.After)
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
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeFHIRDatastoreValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFHIRDatastore(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeFHIRDatastore(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "healthlake",
		OperationName: "DescribeFHIRDatastore",
	}
}
