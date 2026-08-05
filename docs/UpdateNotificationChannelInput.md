# UpdateNotificationChannelInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to **map[string]interface{}** | Updated type-specific configuration. Omit to leave unchanged. | [optional] 
**Name** | Pointer to **string** | Updated display name. Omit to leave unchanged. | [optional] 

## Methods

### NewUpdateNotificationChannelInput

`func NewUpdateNotificationChannelInput() *UpdateNotificationChannelInput`

NewUpdateNotificationChannelInput instantiates a new UpdateNotificationChannelInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNotificationChannelInputWithDefaults

`func NewUpdateNotificationChannelInputWithDefaults() *UpdateNotificationChannelInput`

NewUpdateNotificationChannelInputWithDefaults instantiates a new UpdateNotificationChannelInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *UpdateNotificationChannelInput) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateNotificationChannelInput) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateNotificationChannelInput) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateNotificationChannelInput) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetName

`func (o *UpdateNotificationChannelInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNotificationChannelInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNotificationChannelInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateNotificationChannelInput) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


