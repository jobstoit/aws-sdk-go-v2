// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Grants a temporary extension to the remaining quota of a usage plan associated
// with a specified API key.
func (c *Client) UpdateUsage(ctx context.Context, params *UpdateUsageInput, optFns ...func(*Options)) (*UpdateUsageOutput, error) {
	stack := middleware.NewStack("UpdateUsage", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateUsageMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpUpdateUsageValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateUsage(options.Region), middleware.Before)
	addAcceptHeader(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "UpdateUsage",
			Err:           err,
		}
	}
	out := result.(*UpdateUsageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The PATCH request to grant a temporary extension to the remaining quota of a
// usage plan associated with a specified API key.
type UpdateUsageInput struct {
	// [Required] The Id of the usage plan associated with the usage data.
	UsagePlanId      *string
	TemplateSkipList []*string
	Title            *string
	// [Required] The identifier of the API key associated with the usage plan in which
	// a temporary extension is granted to the remaining quota.
	KeyId    *string
	Name     *string
	Template *bool
	// A list of update operations to be applied to the specified resource and in the
	// order specified in this list.
	PatchOperations []*types.PatchOperation
}

// Represents the usage data of a usage plan. Create and Use Usage Plans
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-api-usage-plans.html),
// Manage Usage in a Usage Plan
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-create-usage-plans-with-console.html#api-gateway-usage-plan-manage-usage)
type UpdateUsageOutput struct {
	// The ending date of the usage data.
	EndDate *string
	// The plan Id associated with this usage data.
	UsagePlanId *string
	// The starting date of the usage data.
	StartDate *string
	// The usage data, as daily logs of used and remaining quotas, over the specified
	// time interval indexed over the API keys in a usage plan. For example, {...,
	// "values" : { "{api_key}" : [ [0, 100], [10, 90], [100, 10]]}, where {api_key}
	// stands for an API key value and the daily log entry is of the format [used
	// quota, remaining quota].
	Items map[string][][]*int64
	// The current pagination position in the paged result set.
	Position *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateUsageMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateUsage{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateUsage{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateUsage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "UpdateUsage",
	}
}
