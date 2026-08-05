# NotificationChannelResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to **map[string]interface{}** | Type-specific configuration. Sensitive values may be redacted in responses. | [optional] 
**CreatedAt** | Pointer to **string** | When the channel was created (RFC3339). | [optional] 
**Id** | Pointer to **string** | Channel record ID. | [optional] 
**IsVerified** | Pointer to **bool** | Whether the channel has completed verification (email link or webhook test). | [optional] 
**Name** | Pointer to **string** | Display name in notification settings. | [optional] 
**Type** | Pointer to **string** | Channel type (for example email or webhook). | [optional] 

## Methods

### NewNotificationChannelResponse

`func NewNotificationChannelResponse() *NotificationChannelResponse`

NewNotificationChannelResponse instantiates a new NotificationChannelResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationChannelResponseWithDefaults

`func NewNotificationChannelResponseWithDefaults() *NotificationChannelResponse`

NewNotificationChannelResponseWithDefaults instantiates a new NotificationChannelResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *NotificationChannelResponse) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *NotificationChannelResponse) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *NotificationChannelResponse) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *NotificationChannelResponse) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NotificationChannelResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NotificationChannelResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NotificationChannelResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NotificationChannelResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *NotificationChannelResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationChannelResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationChannelResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationChannelResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsVerified

`func (o *NotificationChannelResponse) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *NotificationChannelResponse) GetIsVerifiedOk() (*bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerified

`func (o *NotificationChannelResponse) SetIsVerified(v bool)`

SetIsVerified sets IsVerified field to given value.

### HasIsVerified

`func (o *NotificationChannelResponse) HasIsVerified() bool`

HasIsVerified returns a boolean if a field has been set.

### GetName

`func (o *NotificationChannelResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationChannelResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationChannelResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NotificationChannelResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *NotificationChannelResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationChannelResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationChannelResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NotificationChannelResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


