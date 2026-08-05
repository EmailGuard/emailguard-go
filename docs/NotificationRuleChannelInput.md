# NotificationRuleChannelInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelId** | Pointer to **string** | UUID of a team notification channel. Use this or predefined, not both. | [optional] 
**Predefined** | Pointer to **string** | Built-in channel identifier when not using a custom channel. | [optional] 
**TargetUserId** | Pointer to **string** | Restrict delivery to a specific team member&#39;s user ID. | [optional] 

## Methods

### NewNotificationRuleChannelInput

`func NewNotificationRuleChannelInput() *NotificationRuleChannelInput`

NewNotificationRuleChannelInput instantiates a new NotificationRuleChannelInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationRuleChannelInputWithDefaults

`func NewNotificationRuleChannelInputWithDefaults() *NotificationRuleChannelInput`

NewNotificationRuleChannelInputWithDefaults instantiates a new NotificationRuleChannelInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelId

`func (o *NotificationRuleChannelInput) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *NotificationRuleChannelInput) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *NotificationRuleChannelInput) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *NotificationRuleChannelInput) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetPredefined

`func (o *NotificationRuleChannelInput) GetPredefined() string`

GetPredefined returns the Predefined field if non-nil, zero value otherwise.

### GetPredefinedOk

`func (o *NotificationRuleChannelInput) GetPredefinedOk() (*string, bool)`

GetPredefinedOk returns a tuple with the Predefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredefined

`func (o *NotificationRuleChannelInput) SetPredefined(v string)`

SetPredefined sets Predefined field to given value.

### HasPredefined

`func (o *NotificationRuleChannelInput) HasPredefined() bool`

HasPredefined returns a boolean if a field has been set.

### GetTargetUserId

`func (o *NotificationRuleChannelInput) GetTargetUserId() string`

GetTargetUserId returns the TargetUserId field if non-nil, zero value otherwise.

### GetTargetUserIdOk

`func (o *NotificationRuleChannelInput) GetTargetUserIdOk() (*string, bool)`

GetTargetUserIdOk returns a tuple with the TargetUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUserId

`func (o *NotificationRuleChannelInput) SetTargetUserId(v string)`

SetTargetUserId sets TargetUserId field to given value.

### HasTargetUserId

`func (o *NotificationRuleChannelInput) HasTargetUserId() bool`

HasTargetUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


