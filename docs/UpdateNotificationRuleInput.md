# UpdateNotificationRuleInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channels** | Pointer to [**[]NotificationRuleChannelInput**](NotificationRuleChannelInput.md) | Replaces all channel targets when provided. | [optional] 
**Conditions** | Pointer to **map[string]interface{}** | Updated conditions. Omit to leave unchanged. | [optional] 
**Description** | Pointer to **string** | Updated description. Omit to leave unchanged. | [optional] 
**EventTypes** | Pointer to **[]string** | Updated event types. Omit to leave unchanged. | [optional] 
**IsEnabled** | Pointer to **bool** | Enable or disable delivery for this rule. | [optional] 
**Name** | Pointer to **string** | Updated display name. Omit to leave unchanged. | [optional] 

## Methods

### NewUpdateNotificationRuleInput

`func NewUpdateNotificationRuleInput() *UpdateNotificationRuleInput`

NewUpdateNotificationRuleInput instantiates a new UpdateNotificationRuleInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNotificationRuleInputWithDefaults

`func NewUpdateNotificationRuleInputWithDefaults() *UpdateNotificationRuleInput`

NewUpdateNotificationRuleInputWithDefaults instantiates a new UpdateNotificationRuleInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannels

`func (o *UpdateNotificationRuleInput) GetChannels() []NotificationRuleChannelInput`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *UpdateNotificationRuleInput) GetChannelsOk() (*[]NotificationRuleChannelInput, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *UpdateNotificationRuleInput) SetChannels(v []NotificationRuleChannelInput)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *UpdateNotificationRuleInput) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetConditions

`func (o *UpdateNotificationRuleInput) GetConditions() map[string]interface{}`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *UpdateNotificationRuleInput) GetConditionsOk() (*map[string]interface{}, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *UpdateNotificationRuleInput) SetConditions(v map[string]interface{})`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *UpdateNotificationRuleInput) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateNotificationRuleInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateNotificationRuleInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateNotificationRuleInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateNotificationRuleInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEventTypes

`func (o *UpdateNotificationRuleInput) GetEventTypes() []string`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *UpdateNotificationRuleInput) GetEventTypesOk() (*[]string, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *UpdateNotificationRuleInput) SetEventTypes(v []string)`

SetEventTypes sets EventTypes field to given value.

### HasEventTypes

`func (o *UpdateNotificationRuleInput) HasEventTypes() bool`

HasEventTypes returns a boolean if a field has been set.

### GetIsEnabled

`func (o *UpdateNotificationRuleInput) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *UpdateNotificationRuleInput) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *UpdateNotificationRuleInput) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *UpdateNotificationRuleInput) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetName

`func (o *UpdateNotificationRuleInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNotificationRuleInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNotificationRuleInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateNotificationRuleInput) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


