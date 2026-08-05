# IdDataStruct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | UUID of the created or referenced record. | [optional] 

## Methods

### NewIdDataStruct

`func NewIdDataStruct() *IdDataStruct`

NewIdDataStruct instantiates a new IdDataStruct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdDataStructWithDefaults

`func NewIdDataStructWithDefaults() *IdDataStruct`

NewIdDataStructWithDefaults instantiates a new IdDataStruct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IdDataStruct) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdDataStruct) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdDataStruct) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdDataStruct) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


