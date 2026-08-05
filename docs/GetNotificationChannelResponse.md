# GetNotificationChannelResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Machine-readable status code (for example OK or INVALID_BODY). | 
**Data** | Pointer to [**NotificationChannelResponse**](NotificationChannelResponse.md) |  | [optional] 
**Message** | Pointer to **string** | Human-readable detail when the code is not OK. | [optional] 

## Methods

### NewGetNotificationChannelResponse

`func NewGetNotificationChannelResponse(code string, ) *GetNotificationChannelResponse`

NewGetNotificationChannelResponse instantiates a new GetNotificationChannelResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNotificationChannelResponseWithDefaults

`func NewGetNotificationChannelResponseWithDefaults() *GetNotificationChannelResponse`

NewGetNotificationChannelResponseWithDefaults instantiates a new GetNotificationChannelResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GetNotificationChannelResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetNotificationChannelResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetNotificationChannelResponse) SetCode(v string)`

SetCode sets Code field to given value.


### GetData

`func (o *GetNotificationChannelResponse) GetData() NotificationChannelResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetNotificationChannelResponse) GetDataOk() (*NotificationChannelResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetNotificationChannelResponse) SetData(v NotificationChannelResponse)`

SetData sets Data field to given value.

### HasData

`func (o *GetNotificationChannelResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *GetNotificationChannelResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetNotificationChannelResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetNotificationChannelResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetNotificationChannelResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


