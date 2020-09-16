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

// Creates a usage plan with the throttle and quota limits, as well as the
// associated API stages, specified in the payload.
func (c *Client) CreateUsagePlan(ctx context.Context, params *CreateUsagePlanInput, optFns ...func(*Options)) (*CreateUsagePlanOutput, error) {
	stack := middleware.NewStack("CreateUsagePlan", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateUsagePlanMiddlewares(stack)
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
	addOpCreateUsagePlanValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateUsagePlan(options.Region), middleware.Before)
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
			OperationName: "CreateUsagePlan",
			Err:           err,
		}
	}
	out := result.(*CreateUsagePlanOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The POST request to create a usage plan with the name, description, throttle
// limits and quota limits, as well as the associated API stages, specified in the
// payload.
type CreateUsagePlanInput struct {
	// [Required] The name of the usage plan.
	Name             *string
	Title            *string
	Template         *bool
	TemplateSkipList []*string
	// The throttling limits of the usage plan.
	Throttle *types.ThrottleSettings
	// The key-value map of strings. The valid character set is [a-zA-Z+-=._:/]. The
	// tag key can be up to 128 characters and must not start with aws:. The tag value
	// can be up to 256 characters.
	Tags map[string]*string
	// The associated API stages of the usage plan.
	ApiStages []*types.ApiStage
	// The description of the usage plan.
	Description *string
	// The quota of the usage plan.
	Quota *types.QuotaSettings
}

// Represents a usage plan than can specify who can assess associated API stages
// with specified request limits and quotas. In a usage plan, you associate an API
// by specifying the API's Id and a stage name of the specified API. You add plan
// customers by adding API keys to the plan. Create and Use Usage Plans
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-api-usage-plans.html)
type CreateUsagePlanOutput struct {
	// The identifier of a UsagePlan () resource.
	Id *string
	// The request throttle limits of a usage plan.
	Throttle *types.ThrottleSettings
	// The name of a usage plan.
	Name *string
	// The associated API stages of a usage plan.
	ApiStages []*types.ApiStage
	// The AWS Markeplace product identifier to associate with the usage plan as a SaaS
	// product on AWS Marketplace.
	ProductCode *string
	// The maximum number of permitted requests per a given unit time interval.
	Quota *types.QuotaSettings
	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]*string
	// The description of a usage plan.
	Description *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateUsagePlanMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateUsagePlan{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateUsagePlan{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateUsagePlan(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "CreateUsagePlan",
	}
}
