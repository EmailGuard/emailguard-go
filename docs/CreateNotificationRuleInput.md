# CreateNotificationRuleInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channels** | [**[]NotificationRuleChannelInput**](NotificationRuleChannelInput.md) | One or more channel targets to notify when matching events occur. | 
**Conditions** | Pointer to **map[string]interface{}** | Optional rule conditions (e.g. usage_threshold for USAGE_API_THRESHOLD). | [optional] 
**Description** | Pointer to **string** | Optional longer description of what the rule does. | [optional] 
**EventTypes** | **[]string** | Event type identifiers this rule listens for (at least one required). | 
**IsEnabled** | Pointer to **bool** | When false, the rule is stored but does not deliver events. | [optional] 
**Name** | **string** | Display name for the rule. | 

## Methods

### NewCreateNotificationRuleInput

`func NewCreateNotificationRuleInput(channels []NotificationRuleChannelInput, eventTypes []string, name string, ) *CreateNotificationRuleInput`

NewCreateNotificationRuleInput instantiates a new CreateNotificationRuleInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNotificationRuleInputWithDefaults

`func NewCreateNotificationRuleInputWithDefaults() *CreateNotificationRuleInput`

NewCreateNotificationRuleInputWithDefaults instantiates a new CreateNotificationRuleInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannels

`func (o *CreateNotificationRuleInput) GetChannels() []NotificationRuleChannelInput`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *CreateNotificationRuleInput) GetChannelsOk() (*[]NotificationRuleChannelInput, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *CreateNotificationRuleInput) SetChannels(v []NotificationRuleChannelInput)`

SetChannels sets Channels field to given value.


### GetConditions

`func (o *CreateNotificationRuleInput) GetConditions() map[string]interface{}`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *CreateNotificationRuleInput) GetConditionsOk() (*map[string]interface{}, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *CreateNotificationRuleInput) SetConditions(v map[string]interface{})`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *CreateNotificationRuleInput) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetDescription

`func (o *CreateNotificationRuleInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateNotificationRuleInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateNotificationRuleInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateNotificationRuleInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEventTypes

`func (o *CreateNotificationRuleInput) GetEventTypes() []string`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *CreateNotificationRuleInput) GetEventTypesOk() (*[]string, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *CreateNotificationRuleInput) SetEventTypes(v []string)`

SetEventTypes sets EventTypes field to given value.


### GetIsEnabled

`func (o *CreateNotificationRuleInput) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *CreateNotificationRuleInput) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *CreateNotificationRuleInput) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *CreateNotificationRuleInput) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetName

`func (o *CreateNotificationRuleInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNotificationRuleInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNotificationRuleInput) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


