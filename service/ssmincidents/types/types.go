// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// The action that starts at the beginning of an incident. The response plan
// defines the action.
//
// The following types satisfy this interface:
//
//	ActionMemberSsmAutomation
type Action interface {
	isAction()
}

// The Systems Manager automation document to start as the runbook at the beginning
// of the incident.
type ActionMemberSsmAutomation struct {
	Value SsmAutomation

	noSmithyDocumentSerde
}

func (*ActionMemberSsmAutomation) isAction() {}

// Defines the Amazon Web Services Region and KMS key to add to the replication
// set.
type AddRegionAction struct {

	// The Amazon Web Services Region name to add to the replication set.
	//
	// This member is required.
	RegionName *string

	// The KMS key ID to use to encrypt your replication set.
	SseKmsKeyId *string

	noSmithyDocumentSerde
}

// Use the AttributeValueList to filter by string or integer values.
//
// The following types satisfy this interface:
//
//	AttributeValueListMemberIntegerValues
//	AttributeValueListMemberStringValues
type AttributeValueList interface {
	isAttributeValueList()
}

// The list of integer values that the filter matches.
type AttributeValueListMemberIntegerValues struct {
	Value []int32

	noSmithyDocumentSerde
}

func (*AttributeValueListMemberIntegerValues) isAttributeValueList() {}

// The list of string values that the filter matches.
type AttributeValueListMemberStringValues struct {
	Value []string

	noSmithyDocumentSerde
}

func (*AttributeValueListMemberStringValues) isAttributeValueList() {}

// The Systems Manager automation document process to start as the runbook at the
// beginning of the incident.
//
// The following types satisfy this interface:
//
//	AutomationExecutionMemberSsmExecutionArn
type AutomationExecution interface {
	isAutomationExecution()
}

// The Amazon Resource Name (ARN) of the automation process.
type AutomationExecutionMemberSsmExecutionArn struct {
	Value string

	noSmithyDocumentSerde
}

func (*AutomationExecutionMemberSsmExecutionArn) isAutomationExecution() {}

// The Chatbot chat channel used for collaboration during an incident.
//
// The following types satisfy this interface:
//
//	ChatChannelMemberChatbotSns
//	ChatChannelMemberEmpty
type ChatChannel interface {
	isChatChannel()
}

// The Amazon SNS targets that Chatbot uses to notify the chat channel of updates
// to an incident. You can also make updates to the incident through the chat
// channel by using the Amazon SNS topics.
type ChatChannelMemberChatbotSns struct {
	Value []string

	noSmithyDocumentSerde
}

func (*ChatChannelMemberChatbotSns) isChatChannel() {}

// Used to remove the chat channel from an incident record or response plan.
type ChatChannelMemberEmpty struct {
	Value EmptyChatChannel

	noSmithyDocumentSerde
}

func (*ChatChannelMemberEmpty) isChatChannel() {}

// A conditional statement with which to compare a value, after a timestamp, before
// a timestamp, or equal to a string or integer. If multiple conditions are
// specified, the conditionals become an ANDed statement. If multiple values are
// specified for a conditional, the values are ORd.
//
// The following types satisfy this interface:
//
//	ConditionMemberAfter
//	ConditionMemberBefore
//	ConditionMemberEquals
type Condition interface {
	isCondition()
}

// After the specified timestamp.
type ConditionMemberAfter struct {
	Value time.Time

	noSmithyDocumentSerde
}

func (*ConditionMemberAfter) isCondition() {}

// Before the specified timestamp
type ConditionMemberBefore struct {
	Value time.Time

	noSmithyDocumentSerde
}

func (*ConditionMemberBefore) isCondition() {}

// The value is equal to the provided string or integer.
type ConditionMemberEquals struct {
	Value AttributeValueList

	noSmithyDocumentSerde
}

func (*ConditionMemberEquals) isCondition() {}

// Defines the information about the Amazon Web Services Region you're deleting
// from your replication set.
type DeleteRegionAction struct {

	// The name of the Amazon Web Services Region you're deleting from the replication
	// set.
	//
	// This member is required.
	RegionName *string

	noSmithyDocumentSerde
}

// The dynamic SSM parameter value.
//
// The following types satisfy this interface:
//
//	DynamicSsmParameterValueMemberVariable
type DynamicSsmParameterValue interface {
	isDynamicSsmParameterValue()
}

// Variable dynamic parameters. A parameter value is determined when an incident is
// created.
type DynamicSsmParameterValueMemberVariable struct {
	Value VariableType

	noSmithyDocumentSerde
}

func (*DynamicSsmParameterValueMemberVariable) isDynamicSsmParameterValue() {}

// Used to remove the chat channel from an incident record or response plan.
type EmptyChatChannel struct {
	noSmithyDocumentSerde
}

// An item referenced in a TimelineEvent that is involved in or somehow associated
// with an incident. You can specify an Amazon Resource Name (ARN) for an Amazon
// Web Services resource or a RelatedItem ID.
//
// The following types satisfy this interface:
//
//	EventReferenceMemberRelatedItemId
//	EventReferenceMemberResource
type EventReference interface {
	isEventReference()
}

// The ID of a RelatedItem referenced in a TimelineEvent.
type EventReferenceMemberRelatedItemId struct {
	Value string

	noSmithyDocumentSerde
}

func (*EventReferenceMemberRelatedItemId) isEventReference() {}

// The Amazon Resource Name (ARN) of an Amazon Web Services resource referenced in
// a TimelineEvent.
type EventReferenceMemberResource struct {
	Value string

	noSmithyDocumentSerde
}

func (*EventReferenceMemberResource) isEventReference() {}

// Details about a timeline event during an incident.
type EventSummary struct {

	// The timeline event ID.
	//
	// This member is required.
	EventId *string

	// The time that the event occurred.
	//
	// This member is required.
	EventTime *time.Time

	// The type of event. The timeline event must be Custom Event.
	//
	// This member is required.
	EventType *string

	// The time that the timeline event was last updated.
	//
	// This member is required.
	EventUpdatedTime *time.Time

	// The Amazon Resource Name (ARN) of the incident that the event happened during.
	//
	// This member is required.
	IncidentRecordArn *string

	// A list of references in a TimelineEvent.
	EventReferences []EventReference

	noSmithyDocumentSerde
}

// Filter the selection by using a condition.
type Filter struct {

	// The condition accepts before or after a specified time, equal to a string, or
	// equal to an integer.
	//
	// This member is required.
	Condition Condition

	// The key that you're filtering on.
	//
	// This member is required.
	Key *string

	noSmithyDocumentSerde
}

// The record of the incident that's created when an incident occurs.
type IncidentRecord struct {

	// The Amazon Resource Name (ARN) of the incident record.
	//
	// This member is required.
	Arn *string

	// The time that Incident Manager created the incident record.
	//
	// This member is required.
	CreationTime *time.Time

	// The string Incident Manager uses to prevent duplicate incidents from being
	// created by the same incident in the same account.
	//
	// This member is required.
	DedupeString *string

	// The impact of the incident on customers and applications.
	//
	// This member is required.
	Impact *int32

	// Details about the action that started the incident.
	//
	// This member is required.
	IncidentRecordSource *IncidentRecordSource

	// Who modified the incident most recently.
	//
	// This member is required.
	LastModifiedBy *string

	// The time at which the incident was most recently modified.
	//
	// This member is required.
	LastModifiedTime *time.Time

	// The current status of the incident.
	//
	// This member is required.
	Status IncidentRecordStatus

	// The title of the incident.
	//
	// This member is required.
	Title *string

	// The runbook, or automation document, that's run at the beginning of the
	// incident.
	AutomationExecutions []AutomationExecution

	// The chat channel used for collaboration during an incident.
	ChatChannel ChatChannel

	// The Amazon SNS targets that are notified when updates are made to an incident.
	NotificationTargets []NotificationTargetItem

	// The time at which the incident was resolved. This appears as a timeline event.
	ResolvedTime *time.Time

	// The summary of the incident. The summary is a brief synopsis of what occurred,
	// what's currently happening, and context of the incident.
	Summary *string

	noSmithyDocumentSerde
}

// Details about what created the incident record and when it was created.
type IncidentRecordSource struct {

	// The principal that started the incident.
	//
	// This member is required.
	CreatedBy *string

	// The service that started the incident. This can be manually created from
	// Incident Manager, automatically created using an Amazon CloudWatch alarm, or
	// Amazon EventBridge event.
	//
	// This member is required.
	Source *string

	// The service principal that assumed the role specified in createdBy. If no
	// service principal assumed the role this will be left blank.
	InvokedBy *string

	// The resource that caused the incident to be created.
	ResourceArn *string

	noSmithyDocumentSerde
}

// Details describing an incident record.
type IncidentRecordSummary struct {

	// The Amazon Resource Name (ARN) of the incident.
	//
	// This member is required.
	Arn *string

	// The time the incident was created.
	//
	// This member is required.
	CreationTime *time.Time

	// Defines the impact to customers and applications.
	//
	// This member is required.
	Impact *int32

	// What caused Incident Manager to create the incident.
	//
	// This member is required.
	IncidentRecordSource *IncidentRecordSource

	// The current status of the incident.
	//
	// This member is required.
	Status IncidentRecordStatus

	// The title of the incident. This value is either provided by the response plan or
	// overwritten on creation.
	//
	// This member is required.
	Title *string

	// The time the incident was resolved.
	ResolvedTime *time.Time

	noSmithyDocumentSerde
}

// Basic details used in creating a response plan. The response plan is then used
// to create an incident record.
type IncidentTemplate struct {

	// The impact of the incident on your customers and applications.
	//
	// This member is required.
	Impact *int32

	// The title of the incident.
	//
	// This member is required.
	Title *string

	// Used to stop Incident Manager from creating multiple incident records for the
	// same incident.
	DedupeString *string

	// Tags to assign to the template. When the StartIncident API action is called,
	// Incident Manager assigns the tags specified in the template to the incident.
	IncidentTags map[string]string

	// The Amazon SNS targets that are notified when updates are made to an incident.
	NotificationTargets []NotificationTargetItem

	// The summary of the incident. The summary is a brief synopsis of what occurred,
	// what's currently happening, and context.
	Summary *string

	noSmithyDocumentSerde
}

// Details and type of a related item.
type ItemIdentifier struct {

	// The type of related item.
	//
	// This member is required.
	Type ItemType

	// Details about the related item.
	//
	// This member is required.
	Value ItemValue

	noSmithyDocumentSerde
}

// Describes a related item.
//
// The following types satisfy this interface:
//
//	ItemValueMemberArn
//	ItemValueMemberMetricDefinition
//	ItemValueMemberUrl
type ItemValue interface {
	isItemValue()
}

// The Amazon Resource Name (ARN) of the related item, if the related item is an
// Amazon resource.
type ItemValueMemberArn struct {
	Value string

	noSmithyDocumentSerde
}

func (*ItemValueMemberArn) isItemValue() {}

// The metric definition, if the related item is a metric in Amazon CloudWatch.
type ItemValueMemberMetricDefinition struct {
	Value string

	noSmithyDocumentSerde
}

func (*ItemValueMemberMetricDefinition) isItemValue() {}

// The URL, if the related item is a non-Amazon Web Services resource.
type ItemValueMemberUrl struct {
	Value string

	noSmithyDocumentSerde
}

func (*ItemValueMemberUrl) isItemValue() {}

// The SNS targets that are notified when updates are made to an incident.
//
// The following types satisfy this interface:
//
//	NotificationTargetItemMemberSnsTopicArn
type NotificationTargetItem interface {
	isNotificationTargetItem()
}

// The Amazon Resource Name (ARN) of the SNS topic.
type NotificationTargetItemMemberSnsTopicArn struct {
	Value string

	noSmithyDocumentSerde
}

func (*NotificationTargetItemMemberSnsTopicArn) isNotificationTargetItem() {}

// Information about a Amazon Web Services Region in your replication set.
type RegionInfo struct {

	// The status of the Amazon Web Services Region in the replication set.
	//
	// This member is required.
	Status RegionStatus

	// The most recent date and time that Incident Manager updated the Amazon Web
	// Services Region's status.
	//
	// This member is required.
	StatusUpdateDateTime *time.Time

	// The ID of the KMS key used to encrypt the data in this Amazon Web Services
	// Region.
	SseKmsKeyId *string

	// Information displayed about the status of the Amazon Web Services Region.
	StatusMessage *string

	noSmithyDocumentSerde
}

// The mapping between a Amazon Web Services Region and the key that's used to
// encrypt the data.
type RegionMapInputValue struct {

	// The KMS key used to encrypt the data in your replication set.
	SseKmsKeyId *string

	noSmithyDocumentSerde
}

// Resources that responders use to triage and mitigate the incident.
type RelatedItem struct {

	// Details about the related item.
	//
	// This member is required.
	Identifier *ItemIdentifier

	// A unique ID for a RelatedItem. Don't specify this parameter when you add a
	// RelatedItem by using the UpdateRelatedItems API action.
	GeneratedId *string

	// The title of the related item.
	Title *string

	noSmithyDocumentSerde
}

// Details about the related item you're adding.
//
// The following types satisfy this interface:
//
//	RelatedItemsUpdateMemberItemToAdd
//	RelatedItemsUpdateMemberItemToRemove
type RelatedItemsUpdate interface {
	isRelatedItemsUpdate()
}

// Details about the related item you're adding.
type RelatedItemsUpdateMemberItemToAdd struct {
	Value RelatedItem

	noSmithyDocumentSerde
}

func (*RelatedItemsUpdateMemberItemToAdd) isRelatedItemsUpdate() {}

// Details about the related item you're deleting.
type RelatedItemsUpdateMemberItemToRemove struct {
	Value ItemIdentifier

	noSmithyDocumentSerde
}

func (*RelatedItemsUpdateMemberItemToRemove) isRelatedItemsUpdate() {}

// The set of Amazon Web Services Region that your Incident Manager data will be
// replicated to and the KMS key used to encrypt the data.
type ReplicationSet struct {

	// Details about who created the replication set.
	//
	// This member is required.
	CreatedBy *string

	// When the replication set was created.
	//
	// This member is required.
	CreatedTime *time.Time

	// Determines if the replication set deletion protection is enabled or not. If
	// deletion protection is enabled, you can't delete the last Amazon Web Services
	// Region in the replication set.
	//
	// This member is required.
	DeletionProtected *bool

	// Who last modified the replication set.
	//
	// This member is required.
	LastModifiedBy *string

	// When the replication set was last updated.
	//
	// This member is required.
	LastModifiedTime *time.Time

	// The map between each Amazon Web Services Region in your replication set and the
	// KMS key that's used to encrypt the data in that Region.
	//
	// This member is required.
	RegionMap map[string]RegionInfo

	// The status of the replication set. If the replication set is still pending, you
	// can't use Incident Manager functionality.
	//
	// This member is required.
	Status ReplicationSetStatus

	// The Amazon Resource Name (ARN) of the replication set.
	Arn *string

	noSmithyDocumentSerde
}

// The resource policy that allows Incident Manager to perform actions on resources
// on your behalf.
type ResourcePolicy struct {

	// The JSON blob that describes the policy.
	//
	// This member is required.
	PolicyDocument *string

	// The ID of the resource policy.
	//
	// This member is required.
	PolicyId *string

	// The Amazon Web Services Region that policy allows resources to be used in.
	//
	// This member is required.
	RamResourceShareRegion *string

	noSmithyDocumentSerde
}

// Details of the response plan that are used when creating an incident.
type ResponsePlanSummary struct {

	// The Amazon Resource Name (ARN) of the response plan.
	//
	// This member is required.
	Arn *string

	// The name of the response plan. This can't include spaces.
	//
	// This member is required.
	Name *string

	// The human readable name of the response plan. This can include spaces.
	DisplayName *string

	noSmithyDocumentSerde
}

// Details about the Systems Manager automation document that will be used as a
// runbook during an incident.
type SsmAutomation struct {

	// The automation document's name.
	//
	// This member is required.
	DocumentName *string

	// The Amazon Resource Name (ARN) of the role that the automation document will
	// assume when running commands.
	//
	// This member is required.
	RoleArn *string

	// The automation document's version to use when running.
	DocumentVersion *string

	// The key-value pair to resolve dynamic parameter values when processing a Systems
	// Manager Automation runbook.
	DynamicParameters map[string]DynamicSsmParameterValue

	// The key-value pair parameters to use when running the automation document.
	Parameters map[string][]string

	// The account that the automation document will be run in. This can be in either
	// the management account or an application account.
	TargetAccount SsmTargetAccount

	noSmithyDocumentSerde
}

// A significant event that happened during the incident.
type TimelineEvent struct {

	// A short description of the event.
	//
	// This member is required.
	EventData *string

	// The ID of the timeline event.
	//
	// This member is required.
	EventId *string

	// The time that the event occurred.
	//
	// This member is required.
	EventTime *time.Time

	// The type of event that occurred. Currently Incident Manager supports only the
	// Custom Event type.
	//
	// This member is required.
	EventType *string

	// The time that the timeline event was last updated.
	//
	// This member is required.
	EventUpdatedTime *time.Time

	// The Amazon Resource Name (ARN) of the incident that the event occurred during.
	//
	// This member is required.
	IncidentRecordArn *string

	// A list of references in a TimelineEvent.
	EventReferences []EventReference

	noSmithyDocumentSerde
}

// Details about what caused the incident to be created in Incident Manager.
type TriggerDetails struct {

	// Identifies the service that sourced the event. All events sourced from within
	// Amazon Web Services begin with "aws." Customer-generated events can have any
	// value here, as long as it doesn't begin with "aws." We recommend the use of Java
	// package-name style reverse domain-name strings.
	//
	// This member is required.
	Source *string

	// The time that the incident was detected.
	//
	// This member is required.
	Timestamp *time.Time

	// Raw data passed from either Amazon EventBridge, Amazon CloudWatch, or Incident
	// Manager when an incident is created.
	RawData *string

	// The Amazon Resource Name (ARN) of the source that detected the incident.
	TriggerArn *string

	noSmithyDocumentSerde
}

// Details used when updating the replication set.
//
// The following types satisfy this interface:
//
//	UpdateReplicationSetActionMemberAddRegionAction
//	UpdateReplicationSetActionMemberDeleteRegionAction
type UpdateReplicationSetAction interface {
	isUpdateReplicationSetAction()
}

// Details about the Amazon Web Services Region that you're adding to the
// replication set.
type UpdateReplicationSetActionMemberAddRegionAction struct {
	Value AddRegionAction

	noSmithyDocumentSerde
}

func (*UpdateReplicationSetActionMemberAddRegionAction) isUpdateReplicationSetAction() {}

// Details about the Amazon Web Services Region that you're deleting to the
// replication set.
type UpdateReplicationSetActionMemberDeleteRegionAction struct {
	Value DeleteRegionAction

	noSmithyDocumentSerde
}

func (*UpdateReplicationSetActionMemberDeleteRegionAction) isUpdateReplicationSetAction() {}

type noSmithyDocumentSerde = smithydocument.NoSerde

// UnknownUnionMember is returned when a union member is returned over the wire,
// but has an unknown tag.
type UnknownUnionMember struct {
	Tag   string
	Value []byte

	noSmithyDocumentSerde
}

func (*UnknownUnionMember) isAction()                     {}
func (*UnknownUnionMember) isAttributeValueList()         {}
func (*UnknownUnionMember) isAutomationExecution()        {}
func (*UnknownUnionMember) isChatChannel()                {}
func (*UnknownUnionMember) isCondition()                  {}
func (*UnknownUnionMember) isDynamicSsmParameterValue()   {}
func (*UnknownUnionMember) isEventReference()             {}
func (*UnknownUnionMember) isItemValue()                  {}
func (*UnknownUnionMember) isNotificationTargetItem()     {}
func (*UnknownUnionMember) isRelatedItemsUpdate()         {}
func (*UnknownUnionMember) isUpdateReplicationSetAction() {}
