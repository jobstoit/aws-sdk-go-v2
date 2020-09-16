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

// Gets all the usage plans of the caller's account.
func (c *Client) GetUsagePlans(ctx context.Context, params *GetUsagePlansInput, optFns ...func(*Options)) (*GetUsagePlansOutput, error) {
	stack := middleware.NewStack("GetUsagePlans", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetUsagePlansMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetUsagePlans(options.Region), middleware.Before)
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
			OperationName: "GetUsagePlans",
			Err:           err,
		}
	}
	out := result.(*GetUsagePlansOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The GET request to get all the usage plans of the caller's account.
type GetUsagePlansInput struct {
	Template         *bool
	TemplateSkipList []*string
	Title            *string
	// The current pagination position in the paged result set.
	Position *string
	// The identifier of the API key associated with the usage plans.
	KeyId *string
	Name  *string
	// The maximum number of returned results per page. The default value is 25 and the
	// maximum value is 500.
	Limit *int32
}

// Represents a collection of usage plans for an AWS account. Create and Use Usage
// Plans
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-api-usage-plans.html)
type GetUsagePlansOutput struct {
	// The current page of elements from this collection.
	Items []*types.UsagePlan
	// The current pagination position in the paged result set.
	Position *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetUsagePlansMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetUsagePlans{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetUsagePlans{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetUsagePlans(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetUsagePlans",
	}
}
