// Code generated by smithy-go-codegen DO NOT EDIT.

package healthlake

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/healthlake/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Begins a FHIR Import job.
func (c *Client) StartFHIRImportJob(ctx context.Context, params *StartFHIRImportJobInput, optFns ...func(*Options)) (*StartFHIRImportJobOutput, error) {
	if params == nil {
		params = &StartFHIRImportJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartFHIRImportJob", params, optFns, c.addOperationStartFHIRImportJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartFHIRImportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartFHIRImportJobInput struct {

	// Optional user provided token used for ensuring idempotency.
	//
	// This member is required.
	ClientToken *string

	// The Amazon Resource Name (ARN) that gives AWS HealthLake access permission.
	//
	// This member is required.
	DataAccessRoleArn *string

	// The AWS-generated data store ID.
	//
	// This member is required.
	DatastoreId *string

	// The input properties of the FHIR Import job in the StartFHIRImport job request.
	//
	// This member is required.
	InputDataConfig types.InputDataConfig

	// The output data configuration that was supplied when the export job was created.
	//
	// This member is required.
	JobOutputDataConfig types.OutputDataConfig

	// The name of the FHIR Import job in the StartFHIRImport job request.
	JobName *string

	noSmithyDocumentSerde
}

type StartFHIRImportJobOutput struct {

	// The AWS-generated job ID.
	//
	// This member is required.
	JobId *string

	// The status of an import job.
	//
	// This member is required.
	JobStatus types.JobStatus

	// The AWS-generated data store ID.
	DatastoreId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartFHIRImportJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpStartFHIRImportJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpStartFHIRImportJob{}, middleware.After)
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
	if err = addIdempotencyToken_opStartFHIRImportJobMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartFHIRImportJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartFHIRImportJob(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStartFHIRImportJob struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartFHIRImportJob) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartFHIRImportJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartFHIRImportJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartFHIRImportJobInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opStartFHIRImportJobMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartFHIRImportJob{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartFHIRImportJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "healthlake",
		OperationName: "StartFHIRImportJob",
	}
}
