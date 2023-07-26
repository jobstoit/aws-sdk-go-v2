// Code generated by smithy-go-codegen DO NOT EDIT.

package entityresolution

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/entityresolution/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns the MatchingWorkflow with a given name, if it exists.
func (c *Client) GetMatchingWorkflow(ctx context.Context, params *GetMatchingWorkflowInput, optFns ...func(*Options)) (*GetMatchingWorkflowOutput, error) {
	if params == nil {
		params = &GetMatchingWorkflowInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMatchingWorkflow", params, optFns, c.addOperationGetMatchingWorkflowMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMatchingWorkflowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMatchingWorkflowInput struct {

	// The name of the workflow.
	//
	// This member is required.
	WorkflowName *string

	noSmithyDocumentSerde
}

type GetMatchingWorkflowOutput struct {

	// The timestamp of when the workflow was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// A list of InputSource objects, which have the fields InputSourceARN and
	// SchemaName .
	//
	// This member is required.
	InputSourceConfig []types.InputSource

	// A list of OutputSource objects, each of which contains fields OutputS3Path ,
	// ApplyNormalization , and Output .
	//
	// This member is required.
	OutputSourceConfig []types.OutputSource

	// An object which defines the resolutionType and the ruleBasedProperties
	//
	// This member is required.
	ResolutionTechniques *types.ResolutionTechniques

	// The Amazon Resource Name (ARN) of the IAM role. AWS Entity Resolution assumes
	// this role to access resources on your behalf.
	//
	// This member is required.
	RoleArn *string

	// The timestamp of when the workflow was last updated.
	//
	// This member is required.
	UpdatedAt *time.Time

	// The ARN (Amazon Resource Name) that Entity Resolution generated for the
	// MatchingWorkflow .
	//
	// This member is required.
	WorkflowArn *string

	// The name of the workflow.
	//
	// This member is required.
	WorkflowName *string

	// A description of the workflow.
	Description *string

	// An object which defines an incremental run type and has only incrementalRunType
	// as a field.
	IncrementalRunConfig *types.IncrementalRunConfig

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMatchingWorkflowMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetMatchingWorkflow{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetMatchingWorkflow{}, middleware.After)
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
	if err = addOpGetMatchingWorkflowValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMatchingWorkflow(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetMatchingWorkflow(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "entityresolution",
		OperationName: "GetMatchingWorkflow",
	}
}
