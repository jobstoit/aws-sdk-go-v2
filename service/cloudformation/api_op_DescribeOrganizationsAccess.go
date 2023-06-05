// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves information about the account's OrganizationAccess status. This API
// can be called either by the management account or the delegated administrator by
// using the CallAs parameter. This API can also be called without the CallAs
// parameter by the management account.
func (c *Client) DescribeOrganizationsAccess(ctx context.Context, params *DescribeOrganizationsAccessInput, optFns ...func(*Options)) (*DescribeOrganizationsAccessOutput, error) {
	if params == nil {
		params = &DescribeOrganizationsAccessInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeOrganizationsAccess", params, optFns, c.addOperationDescribeOrganizationsAccessMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeOrganizationsAccessOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeOrganizationsAccessInput struct {

	// [Service-managed permissions] Specifies whether you are acting as an account
	// administrator in the organization's management account or as a delegated
	// administrator in a member account. By default, SELF is specified.
	//   - If you are signed in to the management account, specify SELF .
	//   - If you are signed in to a delegated administrator account, specify
	//   DELEGATED_ADMIN . Your Amazon Web Services account must be registered as a
	//   delegated administrator in the management account. For more information, see
	//   Register a delegated administrator (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-orgs-delegated-admin.html)
	//   in the CloudFormation User Guide.
	CallAs types.CallAs

	noSmithyDocumentSerde
}

type DescribeOrganizationsAccessOutput struct {

	// Presents the status of the OrganizationAccess .
	Status types.OrganizationStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeOrganizationsAccessMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeOrganizationsAccess{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeOrganizationsAccess{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeOrganizationsAccess(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeOrganizationsAccess(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "DescribeOrganizationsAccess",
	}
}
