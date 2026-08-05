# BillingUsageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Machine-readable status code (for example OK or INVALID_BODY). | 
**Data** | Pointer to [**[]BillingUsageItem**](BillingUsageItem.md) |  | [optional] 
**Message** | Pointer to **string** | Human-readable detail when the code is not OK. | [optional] 

## Methods

### NewBillingUsageResponse

`func NewBillingUsageResponse(code string, ) *BillingUsageResponse`

NewBillingUsageResponse instantiates a new BillingUsageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingUsageResponseWithDefaults

`func NewBillingUsageResponseWithDefaults() *BillingUsageResponse`

NewBillingUsageResponseWithDefaults instantiates a new BillingUsageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *BillingUsageResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *BillingUsageResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *BillingUsageResponse) SetCode(v string)`

SetCode sets Code field to given value.


### GetData

`func (o *BillingUsageResponse) GetData() []BillingUsageItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BillingUsageResponse) GetDataOk() (*[]BillingUsageItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BillingUsageResponse) SetData(v []BillingUsageItem)`

SetData sets Data field to given value.

### HasData

`func (o *BillingUsageResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *BillingUsageResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BillingUsageResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BillingUsageResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BillingUsageResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


