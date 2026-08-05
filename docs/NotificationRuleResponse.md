# NotificationRuleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channels** | Pointer to [**[]NotificationRuleChannelResponse**](NotificationRuleChannelResponse.md) | Channel targets notified when the rule matches. | [optional] 
**Conditions** | Pointer to **map[string]interface{}** | Optional conditions (e.g. usage threshold percents). | [optional] 
**CreatedAt** | Pointer to **string** | When the rule was created (RFC3339). | [optional] 
**Description** | Pointer to **string** | Optional longer description. | [optional] 
**EventTypes** | Pointer to **[]string** | Event types that trigger this rule. | [optional] 
**Id** | Pointer to **string** | Rule record ID. | [optional] 
**IsEnabled** | Pointer to **bool** | Whether the rule is actively delivering events. | [optional] 
**Name** | Pointer to **string** | Display name for the rule. | [optional] 

## Methods

### NewNotificationRuleResponse

`func NewNotificationRuleResponse() *NotificationRuleResponse`

NewNotificationRuleResponse instantiates a new NotificationRuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationRuleResponseWithDefaults

`func NewNotificationRuleResponseWithDefaults() *NotificationRuleResponse`

NewNotificationRuleResponseWithDefaults instantiates a new NotificationRuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannels

`func (o *NotificationRuleResponse) GetChannels() []NotificationRuleChannelResponse`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *NotificationRuleResponse) GetChannelsOk() (*[]NotificationRuleChannelResponse, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *NotificationRuleResponse) SetChannels(v []NotificationRuleChannelResponse)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *NotificationRuleResponse) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetConditions

`func (o *NotificationRuleResponse) GetConditions() map[string]interface{}`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *NotificationRuleResponse) GetConditionsOk() (*map[string]interface{}, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *NotificationRuleResponse) SetConditions(v map[string]interface{})`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *NotificationRuleResponse) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NotificationRuleResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NotificationRuleResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NotificationRuleResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NotificationRuleResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *NotificationRuleResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NotificationRuleResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NotificationRuleResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NotificationRuleResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEventTypes

`func (o *NotificationRuleResponse) GetEventTypes() []string`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *NotificationRuleResponse) GetEventTypesOk() (*[]string, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *NotificationRuleResponse) SetEventTypes(v []string)`

SetEventTypes sets EventTypes field to given value.

### HasEventTypes

`func (o *NotificationRuleResponse) HasEventTypes() bool`

HasEventTypes returns a boolean if a field has been set.

### GetId

`func (o *NotificationRuleResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationRuleResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationRuleResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationRuleResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsEnabled

`func (o *NotificationRuleResponse) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *NotificationRuleResponse) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *NotificationRuleResponse) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *NotificationRuleResponse) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetName

`func (o *NotificationRuleResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationRuleResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationRuleResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NotificationRuleResponse) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


