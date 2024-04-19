// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// After you create a solution, you can’t change its configuration. By default,
// all new solutions use automatic training. With automatic training, you incur
// training costs while your solution is active. You can't stop automatic training
// for a solution. To avoid unnecessary costs, make sure to delete the solution
// when you are finished. For information about training costs, see Amazon
// Personalize pricing (https://aws.amazon.com/personalize/pricing/) . Creates the
// configuration for training a model (creating a solution version). This
// configuration includes the recipe to use for model training and optional
// training configuration, such as columns to use in training and feature
// transformation parameters. For more information about configuring a solution,
// see Creating and configuring a solution (https://docs.aws.amazon.com/personalize/latest/dg/customizing-solution-config.html)
// . By default, new solutions use automatic training to create solution versions
// every 7 days. You can change the training frequency. Automatic solution version
// creation starts one hour after the solution is ACTIVE. If you manually create a
// solution version within the hour, the solution skips the first automatic
// training. For more information, see Configuring automatic training (https://docs.aws.amazon.com/personalize/latest/dg/solution-config-auto-training.html)
// . To turn off automatic training, set performAutoTraining to false. If you turn
// off automatic training, you must manually create a solution version by calling
// the CreateSolutionVersion (https://docs.aws.amazon.com/personalize/latest/dg/API_CreateSolutionVersion.html)
// operation. After training starts, you can get the solution version's Amazon
// Resource Name (ARN) with the ListSolutionVersions (https://docs.aws.amazon.com/personalize/latest/dg/API_ListSolutionVersions.html)
// API operation. To get its status, use the DescribeSolutionVersion (https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeSolutionVersion.html)
// . After training completes you can evaluate model accuracy by calling
// GetSolutionMetrics (https://docs.aws.amazon.com/personalize/latest/dg/API_GetSolutionMetrics.html)
// . When you are satisfied with the solution version, you deploy it using
// CreateCampaign (https://docs.aws.amazon.com/personalize/latest/dg/API_CreateCampaign.html)
// . The campaign provides recommendations to a client through the
// GetRecommendations (https://docs.aws.amazon.com/personalize/latest/dg/API_RS_GetRecommendations.html)
// API. Amazon Personalize doesn't support configuring the hpoObjective for
// solution hyperparameter optimization at this time. Status A solution can be in
// one of the following states:
//   - CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
//   - DELETE PENDING > DELETE IN_PROGRESS
//
// To get the status of the solution, call DescribeSolution (https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeSolution.html)
// . If you use manual training, the status must be ACTIVE before you call
// CreateSolutionVersion . Related APIs
//
//   - ListSolutions (https://docs.aws.amazon.com/personalize/latest/dg/API_ListSolutions.html)
//
//   - CreateSolutionVersion (https://docs.aws.amazon.com/personalize/latest/dg/API_CreateSolutionVersion.html)
//
//   - DescribeSolution (https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeSolution.html)
//
//   - DeleteSolution (https://docs.aws.amazon.com/personalize/latest/dg/API_DeleteSolution.html)
//
//   - ListSolutionVersions (https://docs.aws.amazon.com/personalize/latest/dg/API_ListSolutionVersions.html)
//
//   - DescribeSolutionVersion (https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeSolutionVersion.html)
func (c *Client) CreateSolution(ctx context.Context, params *CreateSolutionInput, optFns ...func(*Options)) (*CreateSolutionOutput, error) {
	if params == nil {
		params = &CreateSolutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSolution", params, optFns, c.addOperationCreateSolutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSolutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSolutionInput struct {

	// The Amazon Resource Name (ARN) of the dataset group that provides the training
	// data.
	//
	// This member is required.
	DatasetGroupArn *string

	// The name for the solution.
	//
	// This member is required.
	Name *string

	// When your have multiple event types (using an EVENT_TYPE schema field), this
	// parameter specifies which event type (for example, 'click' or 'like') is used
	// for training the model. If you do not provide an eventType , Amazon Personalize
	// will use all interactions for training with equal weight regardless of type.
	EventType *string

	// We don't recommend enabling automated machine learning. Instead, match your use
	// case to the available Amazon Personalize recipes. For more information, see
	// Choosing a recipe (https://docs.aws.amazon.com/personalize/latest/dg/working-with-predefined-recipes.html)
	// . Whether to perform automated machine learning (AutoML). The default is false .
	// For this case, you must specify recipeArn . When set to true , Amazon
	// Personalize analyzes your training data and selects the optimal
	// USER_PERSONALIZATION recipe and hyperparameters. In this case, you must omit
	// recipeArn . Amazon Personalize determines the optimal recipe by running tests
	// with different values for the hyperparameters. AutoML lengthens the training
	// process as compared to selecting a specific recipe.
	PerformAutoML bool

	// Whether the solution uses automatic training to create new solution versions
	// (trained models). The default is True and the solution automatically creates
	// new solution versions every 7 days. You can change the training frequency by
	// specifying a schedulingExpression in the AutoTrainingConfig as part of solution
	// configuration. For more information about automatic training, see Configuring
	// automatic training (https://docs.aws.amazon.com/personalize/latest/dg/solution-config-auto-training.html)
	// . Automatic solution version creation starts one hour after the solution is
	// ACTIVE. If you manually create a solution version within the hour, the solution
	// skips the first automatic training. After training starts, you can get the
	// solution version's Amazon Resource Name (ARN) with the ListSolutionVersions (https://docs.aws.amazon.com/personalize/latest/dg/API_ListSolutionVersions.html)
	// API operation. To get its status, use the DescribeSolutionVersion (https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeSolutionVersion.html)
	// .
	PerformAutoTraining *bool

	// Whether to perform hyperparameter optimization (HPO) on the specified or
	// selected recipe. The default is false . When performing AutoML, this parameter
	// is always true and you should not set it to false .
	PerformHPO *bool

	// The Amazon Resource Name (ARN) of the recipe to use for model training. This is
	// required when performAutoML is false. For information about different Amazon
	// Personalize recipes and their ARNs, see Choosing a recipe (https://docs.aws.amazon.com/personalize/latest/dg/working-with-predefined-recipes.html)
	// .
	RecipeArn *string

	// The configuration to use with the solution. When performAutoML is set to true,
	// Amazon Personalize only evaluates the autoMLConfig section of the solution
	// configuration. Amazon Personalize doesn't support configuring the hpoObjective
	// at this time.
	SolutionConfig *types.SolutionConfig

	// A list of tags (https://docs.aws.amazon.com/personalize/latest/dg/tagging-resources.html)
	// to apply to the solution.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateSolutionOutput struct {

	// The ARN of the solution.
	SolutionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSolutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateSolution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateSolution{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateSolution"); err != nil {
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
	if err = addOpCreateSolutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSolution(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSolution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateSolution",
	}
}
